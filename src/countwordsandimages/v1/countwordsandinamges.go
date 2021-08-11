package ch5

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

//发送一个HTTP GET请求，并且获取文档的字数与图片数量
func CountWordsAndImages(url string) (words, images int, err error) {
	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}
	res, err := http.Get(url)
	if err != nil {
		return
	}

	doc, err := html.Parse(res.Body)
	res.Body.Close()

	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}

	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n.Type == html.TextNode {
		words++
	} else if n.Type == html.ElementNode && n.Data == "img" {
		images++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		swords, simages := countWordsAndImages(c)
		words += swords
		images += simages
	}
	return
}
