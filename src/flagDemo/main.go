/*
 * @Descripttion:
 * @version: v0.1
 * @Author: Elon C
 * @Date: 2021-07-15 21:10:49
 * @LastEditors: Elon C
 * @LastEditTime: 2021-07-15 21:54:13
 * @FilePath: \Golang_Starting\flagDemo\main.go
 */
package main

import (
	"flag"
	"fmt"
	"os"
)

var q string

func main() {
	flag.StringVar(&q, "q", "", "repo:xxx [is:open]")
	fmt.Println(os.Args[0])
	// flag.CommandLine = flag.NewFlagSet(os.Args[1:], flag.ExitOnError)
	flag.Parse()
	fmt.Println(flag.Arg(1))
	fmt.Println(flag.NFlag())
	fmt.Println(q)
}
