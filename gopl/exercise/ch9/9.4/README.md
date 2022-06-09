# 9.4 
使用通道构造一个把任意多个goroutine串联在一起的流水线程序。在内存耗尽之前能创建的最大流水线级数是多少？一个值穿过整个流水线需要多久？

以下测试均在windows10下进行，可用物理内存4.76GB，可用虚拟内存9.00GB
![](系统配置1.png)
![](系统配置2.png)

# 解答一
在主goroutine中创建一个缓存为1的通道，然后用这个通道依次串联两个goroutine
```go
func main() {
	g1 := make(chan struct{}, 1)
	g1 <- struct{}{}
	for {
		go func(in chan struct{}) {
			v := <-in
			in <- v
		}(g1)
	}
}
```
由崩溃转储信息可知最大goroutine编号为17532524