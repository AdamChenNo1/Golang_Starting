package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

// nodecount 统计HTML文档树内所有的元素个数

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline：%v\n", err)
		os.Exit(1)
	}
	count := nodecount(0, doc)
	fmt.Printf("Count of nodes = %d\n", count)
}

func nodecount(count int, n *html.Node) int {
	if n.Type == html.ElementNode {
		count++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		count = nodecount(count, c)
	}
	return count
}
