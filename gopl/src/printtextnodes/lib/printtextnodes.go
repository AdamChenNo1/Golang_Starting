package ch5

import (
	"golang.org/x/net/html"
)

func PrintTextNodes(texts []string, n *html.Node) []string {
	if n.Type == html.TextNode && n.Parent.Data != "script" && n.Parent.Data != "style" {
		texts = append(texts, n.Data)
		// fmt.Printf("%v", texts)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		texts = PrintTextNodes(texts, c)
	}
	return texts
}
