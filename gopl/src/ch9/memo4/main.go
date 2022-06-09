/*
 * File: \src\ch9\memo1\main.go                                                *
 * Project: Golang_Starting                                                    *
 * Created At: Friday, 2022/05/20 , 13:09:08                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Saturday, 2022/05/21 , 01:06:51                              *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package main

import (
	"fmt"
	"go_start/src/ch9/memo4/memo"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

var n sync.WaitGroup

func main() {
	m := memo.New(httpGetBody)

	for url := range incomingURLs() {
		n.Add(1)
		go func(url string) {
			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				log.Print(err)
			}
			fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
			n.Done()
		}(url)
	}
	n.Wait()
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
