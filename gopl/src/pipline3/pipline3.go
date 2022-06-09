/*
 * File: \src\pipline3\pipline3.go                                             *
 * Project: Golang_Starting                                                    *
 * Created At: Thursday, 2022/05/19 , 12:40:54                                 *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Thursday, 2022/05/19 , 12:51:55                              *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package main

import "fmt"

func main() {
	naturals := make(chan int)

	squares := make(chan int)

	//counter
	go func(out chan<- int) {
		for x := 0; x < 100; x++ {
			out <- x
		}
		close(out)
	}(naturals)

	//squarer
	go func(in <-chan int, out chan<- int) {
		for x := range in {
			out <- x * x
		}
		close(out)
	}(naturals, squares)

	//printer
	for x := range squares {
		fmt.Println(x)
	}
}
