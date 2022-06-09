/*
 * File: \src\ch5\outline\v1\main\main.go                                      *
 * Project: Golang_Starting                                                    *
 * Created At: Tuesday, 2021/08/10 , 13:29:02                                  *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Thursday, 2022/05/19 , 14:42:10                              *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

// outline 输出HTML文本节点树的结构
func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline:%v\n", err)
		os.Exit(1)
	}
	outline(nil, doc)
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // 把标签压入栈中
		fmt.Println(stack)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}
