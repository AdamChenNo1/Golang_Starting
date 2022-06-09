/*
 * File: \go_concurrency\ch3\or_channel\main.go                                *
 * Project: Golang_Starting                                                    *
 * Created At: Thursday, 2022/05/26 , 19:09:51                                 *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Thursday, 2022/05/26 , 19:32:19                              *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	var or func(channels ...<-chan interface{}) <-chan any
	or = func(channels ...<-chan any) <-chan any { // 采用可变的channel切片并返回单个channel
		switch len(channels) {
		case 0:
			return nil
		case 1:
			return channels[0]
		}
		orDone := make(chan any)

		go func() { // 创建了一个goroutine，以便可以不受阻塞地等待channel上的消息。
			defer close(orDone)
			switch len(channels) {
			case 2: // 每一次迭代调用都将至少有两个channel。在这里为需要两个channel的情况采用了约束goroutine数目的优化方法。
				select {
				case <-channels[0]:
				case <-channels[1]:
				}
			default:
				// 在循环到存放所有channel的slice的第三个索引的时候，创建了一个or-channel并从这个channel中选择了一个。
				// 这将形成一个由现有slice的剩余部分组成的树并且返回第一个信号量。
				// 为了使在建立这个树的goroutine退出的时候在树下的goroutine也可以跟着退出，
				// 将这个orDone channel也传递到调用中。
				select {
				case <-channels[0]:
				case <-channels[1]:
				case <-channels[2]:
				case <-or(append(channels[3:], orDone)...):
				}
			}
		}()

		return orDone
	}
	sig := func(after time.Duration) <-chan any {
		// 创建一个channel，当后续时间中指定的时间结束时将关闭该channel
		c := make(chan any)

		go func() {
			defer close(c)
			time.Sleep(after)
		}()

		return c
	}

	start := time.Now() // 大致追踪来自or函数的channel何时开始阻塞。
	<-or(
		sig(2*time.Hour),
		sig(2*time.Minute),
		sig(2*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf("done after %v", time.Since(start)) // 打印读取发生的时间
}
