package main

// findlinks4 输出命令行输入链接返回的HTML文档中的所有链接
import (
	"fmt"
	ch5 "go_start/src/findlinks/v4/lib"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		// urls := []string{"https://www.baidu.com"}
		// for _, url := range urls {
		links, err := ch5.FindLinks(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks4: %v\n", err)
			continue
		}
		for _, link := range links {
			fmt.Println(link)
		}
	}
}
