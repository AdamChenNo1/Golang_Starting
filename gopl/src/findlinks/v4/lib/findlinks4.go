package ch5

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

//visit 将n节点中的每个链接添加到结果中
func visit(links []string, n *html.Node) []string {
	if n == nil {
		return links
	}
	if n.Type == html.ElementNode {
		for _, a := range n.Attr {
			if a.Key == "href" || a.Key == "src" {
				links = append(links, a.Val)
			}
		}
	}
	// 递归访问兄弟节点和子节点
	if n.NextSibling != nil {
		links = visit(links, n.NextSibling)
	}
	links = visit(links, n.FirstChild)
	return links
}

// FindLinks发起一个HTTP的GET请求，解析返回的HTML界面，并返回所有链接
func FindLinks(url string) ([]string, error) {
	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}
	res, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		res.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, res.Status)
	}

	doc, err := html.Parse(res.Body)
	res.Body.Close()

	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	return visit(nil, doc), nil
}
