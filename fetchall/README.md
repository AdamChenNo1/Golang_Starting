# fetchall 并发获取多个URL
## v1 
- main函数在一个goroutine中执行，然后go语句创建额外的goroutine
- main函数使用make创建一个字符串通道，对于每个命令行参数，go语句在第一轮循环中启动一个新的goroutine,它异步调用fetch来使用http.Get获取URL内容
- io.Copy函数读取响应的内容，然后通过写入ioutil.Discard输出流进行丢弃。Copy返回字节数以及出现的任何错误
- 每一个结果返回时，fetch发送一行汇总信息到通道ch
- main中的第二轮循环接收并且输出那些汇总行
- 每一个fetch在通道ch上发送一个值（ch <- expression）,main函数接收它们（<- ch）
  - 当一个goroutine试图在一个通道上进行发送或接收操作时，它会阻塞，直到另一个goroutine试图进行接收或发送操作才传递值，并开始处理两个goroutine
  - 由main来处理所有的输出确保了每个goroutine作为一个整体单元处理，这样就避免了两个goroutine同时完成造成输出交织所带来的风险 
## Todo
v1 测试
练习 1.10 
练习 1.11