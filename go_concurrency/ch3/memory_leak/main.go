/*
 * File: \go_concurrency\ch3\memory_leak\main.go                               *
 * Project: Golang_Starting                                                    *
 * Created At: Thursday, 2022/05/26 , 18:08:42                                 *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Thursday, 2022/05/26 , 19:06:20                              *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// test2()
	// test3()
	test4()
}

func memory_leak() {
	doWork := func(strings <-chan string) <-chan interface{} {
		completed := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited .")
			defer close(completed)
			for s := range strings {
				//做些有趣的操作
				fmt.Println(s)
			}
		}()
		return completed
	}
	doWork(nil)
	// 也许这里有其他的操作需要进行
	fmt.Println("Done.")
}

func test2() {
	doWork := func(done <-chan any, strings <-chan string) <-chan any { // 将完成的channel传递给doWork函数。作为惯例，这个channel是第一个参数。
		terminated := make(chan any)
		go func() {
			defer fmt.Println("doWork exited .")
			defer close(terminated)
			for {
				select {
				case s := <-strings:
					// 做些有趣的操作
					fmt.Println(s)
				case <-done: // 检查done channel是否已经发出信号。如果有的话，从goroutine返回。
					return
				}
			}
		}()
		return terminated
	}
	done := make(chan any)
	terminated := doWork(done, nil)
	go func() { // 创建另一个goroutine，如果超过1s就会取消doWork中产生的goroutine。
		time.Sleep(1 * time.Second)
		fmt.Println("Canceling doWork goroutine...")
		close(done)
	}()

	<-terminated //加入从main goroutine的doWork中产生的goroutine的地方。
	//也许这里有其他的操作需要进行
	fmt.Println("Done.")
}

func test3() {
	newRandStream := func() <-chan int {
		randStream := make(chan int)
		go func() {
			defer fmt.Println("newRandStream closure exited. ") // 在goroutine 成功终止时打印出一条消息。
			defer close(randStream)
			for {
				randStream <- rand.Int()
			}
		}()
		return randStream
	}
	randStream := newRandStream()
	fmt.Println("3 random ints:")
	for i := 1; i <= 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}
}

func test4() {
	newRandStream := func(done chan any) <-chan int {
		randStream := make(chan int)
		go func() {
			defer fmt.Println("newRandStream closure exited. ") // 在goroutine 成功终止时打印出一条消息。
			defer close(randStream)
			for {
				select {
				case randStream <- rand.Int():
				case <-done:
					return
				}
			}
		}()
		return randStream
	}

	done := make(chan any)
	randStream := newRandStream(done)
	fmt.Println("3 random ints:")
	for i := 1; i <= 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}
	close(done)

	time.Sleep(1 * time.Second)
}
