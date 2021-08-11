package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	filename := "C:\\Users\\Elon\\Desktop\\golang\\Golang_Starting\\test\\index.html"

	file, _ := os.Open(filename)
	doc, _ := html.Parse(file)
	print(doc)
}

func print(n *html.Node) {
	switch n.Type {
	case html.TextNode:
		fmt.Printf("TextNode: %s", n.Data)
	case html.CommentNode:
		fmt.Printf("CommentNode: %s", n.Data)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		print(c)
	}
}
