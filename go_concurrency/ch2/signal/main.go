/*
 * File: \go_concurrency\ch2\cond\main.go                                      *
 * Project: Golang_Starting                                                    *
 * Created At: Thursday, 2022/05/26 , 15:54:45                                 *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Thursday, 2022/05/26 , 16:02:37                              *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := sync.NewCond(&sync.Mutex{})    // 使用标准的sync.Mutex作为锁
	queue := make([]interface{}, 0, 10) //创建一个长度为0的切片。因为最终会添加10个项目，所以用10的容量实例化它

	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)
		c.L.Lock()        // 再次进入临界区，以便可以修改与条件相关的数据。
		queue = queue[1:] // 通过将切片的头部重新分配到第二个项目来模拟对一个项目的排队。
		fmt.Println("Removed from queue")
		c.L.Unlock() // 退出条件的临界区，因为已经成功地删除了一个项目。
		c.Signal()   // 让一个正在等待的goroutine知道发生了什么事情。
	}

	for i := 0; i < 10; i++ {
		c.L.Lock() // 通过在条件的锁存器上调用锁来进入临界区

		for len(queue) == 2 { // 检查一个循环中队列的长度。这很重要，因为在这种情况下的信号并不一定意味着是所等待的信号，也可能只是发生了什么。
			c.Wait() // 调用Wait，这将暂停main goroutine直到一个信号的条件已经发送。
		}

		fmt.Println("Adding to queue")
		queue = append(queue, struct{}{})

		go removeFromQueue(1 * time.Second) // 创建了一个新的goroutine，它将在一秒钟后删除一个元素。

		c.L.Unlock() // 退出条件的临界区，因为已经成功地进入了一个项目
	}
}
