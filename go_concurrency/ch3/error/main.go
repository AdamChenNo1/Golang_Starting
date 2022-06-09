/*
 * File: \go_concurrency\ch3\error\main.go                                     *
 * Project: Golang_Starting                                                    *
 * Created At: Thursday, 2022/05/26 , 20:07:27                                 *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Thursday, 2022/05/26 , 20:41:02                              *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package main

import (
	"fmt"
	"net/http"
)

func main() {
	// coping_with_errors_in_way_2()
	coping_with_errors_in_way_3()
}

func coping_with_errors_in_way_1() {
	checkStatus := func(
		done <-chan any,
		urls ...string,
	) <-chan *http.Response {
		responses := make(chan *http.Response)
		go func() {
			defer close(responses)
			for _, url := range urls {
				resp, err := http.Get(url)
				if err != nil {
					fmt.Println(err)
					continue
				}
				select {
				case <-done:
					return
				case responses <- resp:
				}
			}
		}()
		return responses
	}
	done := make(chan any)
	defer close(done)
	urls := []string{"https://www.baidu.com", "https://badhost"}
	for response := range checkStatus(done, urls...) {
		fmt.Printf("Response%v\n", response.Status)
	}
}

func coping_with_errors_in_way_2() {
	type Result struct { // 创建一个包含*http.Response和从goroutine中的循环迭代中可能出现的错误的类型。
		Error    error
		Response *http.Response
	}

	// 返回一个可读取的channel，以检索循环迭代的结果。
	checkStatus := func(done <-chan any, urls ...string) <-chan Result {
		results := make(chan Result)

		go func() {
			defer close(results)

			for _, url := range urls {
				var result Result
				resp, err := http.Get(url)
				result = Result{Error: err, Response: resp} // 创建一个Result实例，并设置错误和响应字段。

				select {
				case <-done:
					return
				case results <- result: // 将结果写入channel
				}
			}
		}()
		return results
	}

	done := make(chan any)
	defer close(done)
	urls := []string{
		"https://www.baidu.com",
		"https://badhost",
	}

	for result := range checkStatus(done, urls...) {
		if result.Error != nil {
			fmt.Printf("error :%v\n", result.Error)
			continue
		}
		fmt.Printf("Response :%v\n", result.Response.Status)
	}
}

func coping_with_errors_in_way_3() {
	type Result struct { // 创建一个包含*http.Response和从goroutine中的循环迭代中可能出现的错误的类型。
		Error    error
		Response *http.Response
	}

	// 返回一个可读取的channel，以检索循环迭代的结果。
	checkStatus := func(done <-chan any, urls ...string) <-chan Result {
		results := make(chan Result)

		go func() {
			defer close(results)

			for _, url := range urls {
				var result Result
				resp, err := http.Get(url)
				result = Result{Error: err, Response: resp} // 创建一个Result实例，并设置错误和响应字段。

				select {
				case <-done:
					return
				case results <- result: // 将结果写入channel
				}
			}
		}()
		return results
	}

	done := make(chan any)
	defer close(done)
	urls := []string{
		"a",
		"https://www.baidu.com",
		"https://badhost",
		"b",
		"c",
	}
	errCount := 0

	for result := range checkStatus(done, urls...) {
		if result.Error != nil {
			fmt.Printf("error :%v\n", result.Error)
			errCount++
			if errCount > 3 {
				fmt.Println("Too many errors, breaking!")
				break
			}
			continue
		}
		fmt.Printf("Response :%v\n", result.Response.Status)
	}
}
