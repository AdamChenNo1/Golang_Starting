/*
 * @Descripttion:
 * @version: v0.1
 * @Author: Elon C
 * @Date: 2021-05-22 11:25:02
 * @LastEditors: Elon C
 * @LastEditTime: 2021-05-24 22:40:31
 * @FilePath: \GoPath\src\dup\v2\main.go
 */

//	dup v2 打印输入中多次出现的行的个数和文本
// 它从 stdin 或指定的文件列表读取
package main

import (
	"bufio"
	"fmt"
	"os"
)

var in *os.File = os.Stdin

func dup() {
	counts := make(map[string]int)

	files := os.Args[1:]
	if len(files) == 0 {
		countLines(in, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup_v2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
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
