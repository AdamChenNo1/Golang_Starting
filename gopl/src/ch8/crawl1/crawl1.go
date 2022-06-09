/*
 * File: \src\ch8\crawl1\crawl1.go                                             *
 * Project: gopl                                                               *
 * Created At: Thursday, 2022/05/19 , 14:07:12                                 *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Wednesday, 2022/06/8 , 23:43:18                              *
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

func crawl(url string) []string {
	fmt.Println(url)

	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}

	return list
}
