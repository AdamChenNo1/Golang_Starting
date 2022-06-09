/*
 * @Descripttion:
 * @version: v0.1
 * @Author: Elon C
 * @Date: 2021-05-19 18:01:35
 * @LastEditors: Elon C
 * @LastEditTime: 2021-05-19 18:13:29
 * @FilePath: \GoPath\src\charcount\main.go
 */
//charcount 计算Unicode字符的个数
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)    //Unicode字符数量
	var utflen [utf8.UTFMax + 1]int //UTF-8编码的长度
	invalid := 0                    //非法UTF-8字符数量

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() //返回rune:解码的字符,nbytes:UTF-8编码中字节的长度,error错误
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount:%v", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 { //输入的是不合法的UTF-8字符
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}

	fmt.Printf("\nrune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Printf("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF characters\n", invalid)
	}
}
