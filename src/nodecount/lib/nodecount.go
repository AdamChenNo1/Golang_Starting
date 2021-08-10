package ch5

import "golang.org/x/net/html"

func NodeCount(count int, n *html.Node) int {
	if n.Type == html.ElementNode {
		count++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		count = NodeCount(count, c)
	}
	return count
}
