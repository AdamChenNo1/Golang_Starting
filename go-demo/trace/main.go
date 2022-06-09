/*
 * File: \test\trace\main.go                                                   *
 * Project: Golang_Starting                                                    *
 * Created At: Sunday, 2022/05/22 , 20:14:11                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Sunday, 2022/05/22 , 23:28:18                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package main

import (
	"os"
	"runtime/trace"
)

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()
	// Your program here
	// trace.Start(os.Stderr)
	// defer trace.Stop()

	ch := make(chan string)

	go func() {
		ch <- "go"
	}()

	<-ch
}
