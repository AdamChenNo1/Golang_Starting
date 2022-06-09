/*
 * File: \go_concurrency\ch2\goroutine_time\goroutine_test.go                  *
 * Project: Golang_Starting                                                    *
 * Created At: Thursday, 2022/05/26 , 15:25:05                                 *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Thursday, 2022/05/26 , 15:39:34                              *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package main

import (
	"sync"
	"testing"
)

func BenchmarkContextSwitch(b *testing.B) {
	var wg sync.WaitGroup
	begin := make(chan struct{})
	c := make(chan struct{})

	var token struct{}
	sender := func() {
		defer wg.Done()
		// 在这里等待，直到被告知开始执行。对上下文切换度量的时候，
		// 不需要考虑设置和启动每个goroutine的成本。
		<-begin

		for i := 0; i < b.N; i++ {
			// 将消息发送到接收器goroutine 。一个struct{}{}被称为一个空结构，
			// 它没有内存占用，因此，只是在发出信号的时候记录时间。
			c <- token // 收到一条信息，但什么也不做
		}
	}

	receiver := func() {
		defer wg.Done()
		<-begin

		for i := 0; i < b.N; i++ {
			<-c
		}
	}

	wg.Add(2)

	go sender()
	go receiver()

	b.StartTimer() // 开始计时
	close(begin)   //告诉两个goroutine开始运行
	wg.Wait()
}
