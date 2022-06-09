/*
 * File: \go_concurrency\ch3\generator\pipline_test.go                         *
 * Project: Golang_Starting                                                    *
 * Created At: Thursday, 2022/05/26 , 23:08:10                                 *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Friday, 2022/05/27 , 18:31:42                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package main

import (
	pipline "go_start/go_concurrency/ch3"
	"testing"
)

func BenchmarkGeneric(b *testing.B) {
	done := make(chan any)
	defer close(done)

	b.ResetTimer()
	for range pipline.ToString(done, pipline.Take(done, pipline.Repeat(done, "a"), b.N)) {
	}
}

func BenchmarkTyped(b *testing.B) {
	repeat := func(done <-chan any, values ...string) <-chan string {
		valueStream := make(chan string)

		go func() {
			defer close(valueStream)
			for {
				for _, v := range values {
					select {
					case <-done:
						return
					case valueStream <- v:
					}
				}
			}
		}()

		return valueStream
	}

	take := func(done <-chan any, valueStream <-chan string, num int) <-chan string {
		takenStream := make(chan string)

		go func() {
			defer close(takenStream)
			for i := 0; i < num; i++ {
				select {
				case <-done:
					return
				case takenStream <- <-valueStream:
				}
			}
		}()

		return takenStream
	}

	done := make(chan any)
	defer close(done)

	b.ResetTimer()
	for range take(done, repeat(done, "a"), b.N) {
	}
}
