/*
 * File: \go_concurrency\ch3\tee_channel\main.go                               *
 * Project: Golang_Starting                                                    *
 * Created At: Friday, 2022/05/27 , 18:19:01                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Friday, 2022/05/27 , 18:44:45                                *
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
)

func main() {
	done := make(chan any)
	defer close(done)

	out1, out2 := pipline.Tee(done, pipline.Take(done, pipline.Repeat(done, 1, 2), 4))

	for val1 := range out1 {
		fmt.Printf("out1: %v,out2: %v\n", val1, <-out2)
	}
}
