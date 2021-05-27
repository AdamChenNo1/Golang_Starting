/*
 * @Descripttion:
 * @version: v0.1
 * @Author: Elon C
 * @Date: 2021-05-27 22:41:21
 * @LastEditors: Elon C
 * @LastEditTime: 2021-05-27 22:45:08
 * @FilePath: \Golang_Starting\server\v1\server.go
 */
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler) //echo请求调用程序
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

// 处理程序回显请求URL r的路径部分
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
