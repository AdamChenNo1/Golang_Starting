/*
 * File: \src\pipline1\pipline1.go                                             *
 * Project: Golang_Starting                                                    *
 * Created At: Thursday, 2022/05/19 , 12:24:01                                 *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Thursday, 2022/05/19 , 12:28:39                              *
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
		for x := 0; ; x++ {
			naturals <- x
		}
	}()

	//squarer
	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
	}()

	//printer
	for {
		fmt.Println(<-squares)
	}
}
