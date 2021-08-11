package main

import (
	"fmt"
	ch2 "go_start/src/outline/v2"
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
