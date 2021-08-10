/*
 * @Descripttion:
 * @version: v0.1
 * @Author: Elon C
 * @Date: 2021-07-15 20:38:18
 * @LastEditors: Elon C
 * @LastEditTime: 2021-07-15 20:42:02
 * @FilePath: \Golang_Starting\sleep\main.go
 */
package main

import (
	"flag"
	"fmt"
	"time"
)

var period = flag.Duration("period", 1*time.Second, "sleep period")

func main() {
	flag.Parse()
	fmt.Printf("Sleeping for %v...", *period)
	time.Sleep(*period)
	fmt.Println()
}
