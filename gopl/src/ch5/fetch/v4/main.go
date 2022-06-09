/*
 * @Descripttion:
 * @version: v0.1
 * @Author: Elon C
 * @Date: 2021-05-26 00:11:22
 * @LastEditors: Elon C
 * @LastEditTime: 2021-05-26 00:52:33
 * @FilePath: \Golang_Starting\fetch\v4\main.go
 */
// fetch v2 获取指定URL的内容，然后不加解析地输出
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url) //产生一个http请求
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		// fmt.Println(resp.Status)
		// b, err := ioutil.ReadAll(resp.Body) //读取整个响应结果并存入b
		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}
