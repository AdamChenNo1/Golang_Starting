/*
 * File: \src\ch5\outline\v2\main\main.go                                      *
 * Project: gopl                                                               *
 * Created At: Wednesday, 2021/08/11 , 19:31:06                                *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Wednesday, 2022/06/8 , 23:40:56                              *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package main

import (
	"fmt"
	ch2 "go_start/gopl/src/ch5/outline/v2"
	"os"

	"golang.org/x/net/html"
)

// outline 输出HTML文本节点树的结构
func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline：%v\n", err)
		os.Exit(1)
	}
	ch2.ForEachNode(doc, ch2.StartElement, ch2.EndElement)
}
