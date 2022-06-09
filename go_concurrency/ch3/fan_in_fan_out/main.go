/*
 * File: \go_concurrency\ch3\fan_in_fan_out\main.go                            *
 * Project: Golang_Starting                                                    *
 * Created At: Friday, 2022/05/27 , 17:21:36                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Friday, 2022/05/27 , 18:33:20                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package main

import (
	"fmt"
	pipline "go_start/go_concurrency/ch3"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func test1() {
	rand := func() any {
		return rand.Intn(50000000)
	}

	done := make(chan any)
	defer close(done)

	start := time.Now()
	randIntStream := toInt(done, pipline.RepeatFn(done, rand))

	fmt.Println("Primes:")

	for prime := range pipline.Take(done, primeFinder(done, randIntStream), 10) {
		fmt.Printf("\t%d\n", prime)
	}

	fmt.Printf("Search took: %v", time.Since(start))
}

func test2() {
	//采用我们的标准done channel来允许goroutine被拆除，然后是一个可变的interface{} channel用来扇入
	fanIn := func(done <-chan any, channels ...<-chan any) <-chan any {
		var wg sync.WaitGroup // 创建了一个sync.WaitGroup，以便可以等待直到所有channel都已耗尽
		multiplexedStream := make(chan any)

		// multiplex 在传递时将从channel中读取，并将读取的值传递到multiplexedStream channel
		multiplex := func(c <-chan any) {
			defer wg.Done()
			for i := range c {
				select {
				case <-done:
					return
				case multiplexedStream <- i:
				}
			}
		}

		// 从所有的channel 里取值
		wg.Add(len(channels))
		for _, c := range channels {
			go multiplex(c)
		}

		// 等待所有的读操作结束
		go func() {
			wg.Wait()
			close(multiplexedStream)
		}()

		return multiplexedStream
	}

	rand := func() any {
		return rand.Intn(50000000)
	}

	done := make(chan any)
	defer close(done)

	start := time.Now()
	randIntStream := toInt(done, pipline.RepeatFn(done, rand))

	numFinders := runtime.NumCPU()
	fmt.Printf("Spinning up %d prime finders.\n", numFinders)
	finders := make([]<-chan any, numFinders)
	fmt.Println("Primes:")

	for i := 0; i < numFinders; i++ {
		finders[i] = primeFinder(done, randIntStream)
	}

	for prime := range pipline.Take(done, fanIn(done, finders...), 10) {
		fmt.Printf("\t%d\n", prime)
	}

	fmt.Printf("Search took: %v", time.Since(start))
}

func primeFinder(done <-chan any, valueStream <-chan int) <-chan any

func toInt(done <-chan any, valueStream <-chan any) <-chan int
