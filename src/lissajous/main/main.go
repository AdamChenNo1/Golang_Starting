/*
 * @Descripttion:
 * @version: v0.1
 * @Author: Elon C
 * @Date: 2021-05-24 23:06:23
 * @LastEditors: Elon C
 * @LastEditTime: 2021-05-28 20:00:17
 * @FilePath: \Golang_Starting\lissajous\main\main.go
 */
//lissajous 产生随机利萨茹图形的GIF动画
package main

import (
	lissajous "gostart/lissajous"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajous.Lissajous(w)
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8080", nil))
		return
	}

	lissajous.Lissajous(os.Stdout)
}
