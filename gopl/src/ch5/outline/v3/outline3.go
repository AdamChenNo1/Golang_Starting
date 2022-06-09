package ch2

import (
	"fmt"
	"strings"

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
		if n.FirstChild == nil {
			fmt.Printf("%*s<%s/", depth*2, "", n.Data)
			for _, a := range n.Attr {
				fmt.Printf(" %s=\"%s\"", a.Key, a.Val)
			}
			fmt.Print(">\n")
		} else {
			fmt.Printf("%*s<%s", depth*2, "", n.Data)
			for _, a := range n.Attr {
				fmt.Printf(" %s=\"%s\"", a.Key, a.Val)
			}
			fmt.Print(">\n")
			depth++
		}
	} else if n.Type == html.TextNode || n.Type == html.CommentNode {
		if text := strings.TrimSpace(n.Data); text != "\n" && text != "" {
			fmt.Printf("%*s%s\n", depth*2, "", text)
		}
	}
}

func EndElement(n *html.Node) {
	if n.Type == html.ElementNode && n.FirstChild != nil {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}
