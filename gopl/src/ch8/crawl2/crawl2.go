/*
 * File: \src\ch8\crawl2\crawl2.go                                             *
 * Project: gopl                                                               *
 * Created At: Thursday, 2022/05/19 , 15:41:46                                 *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Wednesday, 2022/06/8 , 23:43:27                              *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package main

import (
	"fmt"
	"go_start/gopl/src/ch5/links"
	"log"
)

var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)

	tokens <- struct{}{} //obtain tokens
	list, err := links.Extract(url)
	<-tokens //release tokens

	if err != nil {
		log.Print(err)
	}

	return list
}
