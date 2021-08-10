package main

import (
	"fmt"
	"os"

	ch5 "go_start/src/nodecount/lib"

	"golang.org/x/net/html"
)

// nodecount 统计HTML文档树内所有的元素个数

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline：%v\n", err)
		os.Exit(1)
	}
	count := ch5.NodeCount(0, doc)
	fmt.Printf("Count of nodes = %d\n", count)
}
