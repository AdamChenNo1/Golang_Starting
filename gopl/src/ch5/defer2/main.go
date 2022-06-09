/*
 * File: \src\ch5\defer1\main.go                                               *
 * Project: Golang_Starting                                                    *
 * Created At: Saturday, 2022/05/21 , 17:08:29                                 *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Saturday, 2022/05/21 , 17:12:33                              *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	defer printStack()
	f(3)
}

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x) //当x == 0时宕机
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}
