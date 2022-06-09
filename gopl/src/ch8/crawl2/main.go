/*
 * File: \src\ch8\crawl2\main.go                                               *
 * Project: Golang_Starting                                                    *
 * Created At: Thursday, 2022/05/19 , 15:41:31                                 *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Thursday, 2022/05/19 , 15:49:15                              *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package main

import "os"

func main() {
	worklist := make(chan []string)

	var n int

	n++
	go func() {
		worklist <- os.Args[1:]
	}()

	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
			}

			n++

			go func(link string) {
				worklist <- crawl(link)
			}(link)
		}
	}
}
