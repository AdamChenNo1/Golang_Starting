/*
 * File: \go_concurrency\ch2\goroutine_memory\main.go                          *
 * Project: Golang_Starting                                                    *
 * Created At: Thursday, 2022/05/26 , 15:11:25                                 *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Thursday, 2022/05/26 , 15:12:19                              *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	memConsumed := func() uint64 {
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}

	var c <-chan interface{}
	var wg sync.WaitGroup

	//需要一个永远不会退出的goroutine，这样就可以在内存中保留一
	//段时间用于测算。不要担心是如何实现这个目标的，只要了解这个
	//goroutine 不会退出，直到进程结束。
	noop := func() { wg.Done(); <-c }

	//定义了要创建的goroutine 的数量。我们将用大数定律，惭惭地接近一个
	//goroutine 的大小。
	const numGoroutines = 1e4

	wg.Add(numGoroutines)
	before := memConsumed() //测算在创建goroutine之前消耗的内存总量
	for i := numGoroutines; i > 0; i-- {
		go noop()
	}
	wg.Wait()

	after := memConsumed() //测算在创建goroutine之后消耗的内存总量
	fmt.Printf("%.3fkb", float64(after-before)/numGoroutines/1000)
}
