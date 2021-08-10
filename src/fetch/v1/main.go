/*
 * @Descripttion:
 * @version: v0.1
 * @Author: Elon C
 * @Date: 2021-05-26 00:11:22
 * @LastEditors: Elon C
 * @LastEditTime: 2021-05-26 00:31:45
 * @FilePath: \Golang_Starting\fetch\v1\main.go
 */
// fetch v1 获取指定URL的内容，然后不加解析地输出
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url) //产生一个http请求
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body) //读取整个响应结果并存入b
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
