package main

import (
	"fmt"
	ch5 "go_start/src/printtextnodes/lib"
	"os"

	"golang.org/x/net/html"
)

// printtextnodes 输出HTML文档树中所有文本节点的内容，不包括`<script>`和`<style>`元素
func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "printtextnodes：%v\n", err)
		os.Exit(1)
	}

	for _, text := range ch5.PrintTextNodes(nil, doc) {
		fmt.Println(text)
	}
}
