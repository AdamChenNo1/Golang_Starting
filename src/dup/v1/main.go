/*
 * @Descripttion:
 * @version: v0.1
 * @Author: Elon C
 * @Date: 2021-05-22 11:25:02
 * @LastEditors: Elon C
 * @LastEditTime: 2021-05-24 16:06:55
 * @FilePath: \GoPath\src\dup\main.go
 */

//	dup v1 输出标准输入中出现次数大于1的行，前面是次数
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var in io.Reader = os.Stdin

func dup() {
	counts := make(map[string]int)

	input := bufio.NewScanner(in)
	for input.Scan() {
		counts[input.Text()]++
	}

	// 注意：忽略input.Err()中可能的错误
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func main() {
	dup()
}
