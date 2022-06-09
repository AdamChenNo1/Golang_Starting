/*
 * File: \src\ch9\memo1\main.go                                                *
 * Project: gopl                                                               *
 * Created At: Friday, 2022/05/20 , 13:09:08                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Wednesday, 2022/06/8 , 23:49:53                              *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package main

import (
	"fmt"
	"go_start/gopl/src/ch9/memo1/memo"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	m := memo.New(httpGetBody)

	for url := range incomingURLs() {
		start := time.Now()
		value, err := m.Get(url)
		if err != nil {
			log.Print(err)
		}
		fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
	}
}

func httpGetBody(url string) (any, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func incomingURLs() <-chan string {
	ch := make(chan string)
	go func() {
		for _, url := range []string{
			"https://go.dev/",
			"https://pkg.go.dev/",
			"https://go.dev/play/",
			"http://gopl.io",
			"https://go.dev/",
			"https://pkg.go.dev/",
			"https://go.dev/play/",
			"http://gopl.io",
		} {
			ch <- url
		}
		close(ch)
	}()
	return ch
}
