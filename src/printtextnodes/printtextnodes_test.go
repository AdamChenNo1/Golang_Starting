package main

import (
	"fmt"
	ch5 "go_start/src/printtextnodes/lib"
	"os"

	"golang.org/x/net/html"
)

// nodecount 统计HTML文档树内所有的元素个数
func ExamplePrintTextNodes() {
	filename := "index.html"

	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	doc, err := html.Parse(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "PrintTextNodes：%v\n", err)
	}
	for _, text := range ch5.PrintTextNodes(nil, doc) {
		fmt.Println(text)
	}
}

// Hello
// World
