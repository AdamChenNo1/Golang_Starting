/*
 * File: \src\links\links.go                                                   *
 * Project: Golang_Starting                                                    *
 * Created At: Thursday, 2021/08/12 , 10:57:40                                 *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Thursday, 2022/05/19 , 15:20:26                              *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package links

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

// ForEachNode调用pre(x)和post(x)遍历以n为根的树中的每个节点
// pre在子节点被访问前调用
// post在子节点被访问后调用
func ForEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ForEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

// Extract 向给定URL发起HTTP GET请求，解析并返回HTML文档中存在的链接
func Extract(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %v", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()

	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					continue
				}

				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue
				}

				links = append(links, link.String())
			}
		}
	}

	ForEachNode(doc, visitNode, nil)
	return links, nil
}
