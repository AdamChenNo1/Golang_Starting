package ch2

import (
	"fmt"

	"golang.org/x/net/html"
)

// ForEachNode调用pre(x)和post(x)遍历以n为根的树中的每个节点
// pre在子节点被访问前调用
// post在子节点被访问后调用
func ForEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if n != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ForEachNode(c, pre, post)
	}

	if n != nil {
		post(n)
	}
}

var depth int

func StartElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}
}

func EndElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}
