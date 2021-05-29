/*
 * @Descripttion:
 * @version: v0.1
 * @Author: Elon C
 * @Date: 2021-05-28 23:02:24
 * @LastEditors: Elon C
 * @LastEditTime: 2021-05-28 23:05:18
 * @FilePath: \Golang_Starting\echo\v4\echo.go
 */
// echo v4 输出其命令行参数
package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
