/*
 * File: \src\ch8\crawl1\main.go                                               *
 * Project: Golang_Starting                                                    *
 * Created At: Thursday, 2022/05/19 , 14:08:49                                 *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Thursday, 2022/05/19 , 14:12:59                              *
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

	//start from commandline arguments
	go func() {
		worklist <- os.Args[1:]
	}()

	//crawl web concurrently
	seen := make(map[string]bool)

	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true

				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}
