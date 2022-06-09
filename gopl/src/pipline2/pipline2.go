/*
 * File: \src\pipline2\pipline2.go                                             *
 * Project: Golang_Starting                                                    *
 * Created At: Thursday, 2022/05/19 , 12:36:36                                 *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Thursday, 2022/05/19 , 12:39:37                              *
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
	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	//squarer
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	//printer
	for x := range squares {
		fmt.Println(x)
	}
}
