/*
 * File: \go_concurrency\ch3\bridge\main.go                                    *
 * Project: go_concurrency                                                     *
 * Created At: Friday, 2022/05/27 , 19:00:02                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Wednesday, 2022/06/8 , 23:48:27                              *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package main

import (
	"fmt"
	pipline "go_start/go_concurrency/ch3"
)

func main() {
	genVals := func() <-chan <-chan any {
		chanStream := make(chan (<-chan any))
		go func() {
			defer close(chanStream)
			for i := 0; i < 10; i++ {
				stream := make(chan any, 1)
				stream <- i
				close(stream)
				chanStream <- stream
			}
		}()
		return chanStream
	}

	for v := range pipline.Bridge(nil, genVals()) {
		fmt.Printf("%v ", v)
	}
}
