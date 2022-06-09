# 并发非阻塞的缓存

## 函数记忆问题
缓存函数的结果，达到多次调用但只需计算一次的效果

并发安全版本
- 将map变量限制在一个goroutine中
并行请求
- Memo类型包含一个通道requests，通道类型为request，仅保存一个值
- Get函数的调用者通过requests通道向监控者goroutine发送被记忆函数f的参数key，以及一个通道response，结果在准备好通过response通道发回
- Get方法创建了一个响应通道，放在了请求里，然后发送给监控goroutine，接着从通道中读取
- cache变量被限制在监控goroutine，即((*Memo).server)中，监控goroutine从request通道中循环读取，直到该通道被Close方法关闭。对于每个请求，其先查询缓存，如果没找到，则创建并插入一个新的entry
- (*entry).call对于指定键的一次请求，在该键上调用函数f，保存到结果entry中，最后通过关闭通道ready来广播准备完毕状态
- 对同一个键的后续请求会在map中找到已有的entry，然后等待结果准备好，最后通过响应通道把结果发回给调用Get的客户端goroutine
- call和deliver方法都需要在独立的goroutine，以确保监控goroutine能持续处理请求