/*
 * File: \src\ch5\defer1\main.go                                               *
 * Project: Golang_Starting                                                    *
 * Created At: Saturday, 2022/05/21 , 17:08:29                                 *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Saturday, 2022/05/21 , 17:08:39                              *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package main

import "fmt"

func main() {
	f(3)
}

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x) //当x == 0时宕机
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}
