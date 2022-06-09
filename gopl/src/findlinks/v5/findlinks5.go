package ch5

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"

	"go_start/src/links"
)

// breadFirst 对每个worklist调用f
// 并将返回的内容添加到worklist中，对每个元素，最多调用一次f
func BreadFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

// crawl 输出url，解析链接然后返回
func Crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
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
