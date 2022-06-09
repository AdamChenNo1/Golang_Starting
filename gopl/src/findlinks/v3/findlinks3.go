package main

// findlinks2 输出从标注输入中读入的HTML文档中的所有链接
import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks3: %v\n", err)
		os.Exit(1)
	}

	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

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
