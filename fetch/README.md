# fetch 获取指定URL的内容，然后不加解析地输出
# 用法
    fetch [http://]www.example.com [[http://]www.example1.com [http://]www.example2.com]]
    - []内为可选内容
## v1 
- 使用两个标准库net/http和io/ioutil,
- http.Get产生一个http请求，如果没有出错，返回结果存在响应结构resp里面，其中resp的Body域包含服务器端响应的一个可读取数据流。
- 随后ioutil.ReadAll读取整个响应结果并存入b。
- 关闭Body数据流来避免资源泄露，使用Printf将响应输出到标准输出
- 无论哪种错误情况，os.Exit(1)会在进程退出时返回状态码1
## v2 
- 使用io.Copy来代替ioutil.ReadAll来复制相应内容到os.Stdout，这样不需要装下响应数据流的缓冲区  
## v3 
- 假如URL参数缺失协议前缀，则自动添加http://前缀  
## v4 
- 增加输出HTTP的状态码功能  
# Todo
- v1 测试
- v2 测试
- v3 测试