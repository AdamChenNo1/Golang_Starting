/*
 * @Descripttion:
 * @version: v0.1
 * @Author: Elon C
 * @Date: 2021-05-22 11:25:02
 * @LastEditors: Elon C
 * @LastEditTime: 2021-05-24 22:33:47
 * @FilePath: \GoPath\src\dup\v3\main.go
 */

//	dup v3 打印输入中多次出现的行的个数和文本
// 它从指定的文件列表读取输入
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func dup() {
	counts := make(map[string]int)

	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup_v3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}

func main() {
	dup()
}
