/*
 * File: \src\spinner\main.go                                                  *
 * Project: Golang_Starting                                                    *
 * Created At: Wednesday, 2022/05/18 , 00:10:59                                *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Wednesday, 2022/05/18 , 00:20:41                             *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
// spinner 递归计算斐波那契数，在执行期间提供一个可见的指示
package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond) //main函数返回时，所有的goroutine都暴力地终结
	const n = 45
	fibN := fib(n)
	// fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
	fmt.Printf("\bFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}

	return fib(x-1) + fib(x-2)
}
