/*
 * @Descripttion:
 * @version: v0.1
 * @Author: Elon C
 * @Date: 2021-05-22 11:25:02
 * @LastEditors: Elon C
 * @LastEditTime: 2021-05-24 22:56:58
 * @FilePath: \GoPath\src\dup\v4\main.go
 */

//	dup v4 输出出现重复行的文件的名称
// 它从指定的文件列表读取输入
package main

import (
	"bufio"
	"fmt"
	"os"
)

func dup() {
	files := os.Args[1:]
	num := 0
l:
	for _, arg := range files {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup_v4: %v\n", err)
			continue
		}
		defer f.Close()
		counts := make(map[string]int)
		input := bufio.NewScanner(f)

		for input.Scan() {
			text := input.Text()
			counts[text]++
			if counts[text] > 1 {
				num++
				fmt.Printf("%d\t%s\n", num, arg)
				continue l
			}
		}
	}
}

func main() {
	dup()
}
