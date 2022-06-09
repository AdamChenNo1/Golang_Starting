package main

import (
	"fmt"
	ch5 "go_start/src/findlinks/v4/lib"
)

func ExampleFindlinks() {
	url := "http://www.baidu.com"

	links, err := ch5.FindLinks(url)

	if err != nil {
		for _, link := range links {
			fmt.Println(link)
		}
	}
}
