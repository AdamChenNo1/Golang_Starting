/*
 * File: \src\ch12\format\main\main.go                                         *
 * Project: Golang_Starting                                                    *
 * Created At: Monday, 2022/05/23 , 16:45:21                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Monday, 2022/05/23 , 17:12:47                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package main

import (
	"fmt"
	"go_start/src/ch12/format"
	"time"
)

func main() {
	var x int64

	var d time.Duration = 1 * time.Second

	fmt.Println(format.Any(x))
	fmt.Println(format.Any(d))
	fmt.Println(format.Any([]int64{x}))
	fmt.Println(format.Any([]time.Duration{d}))
}
