package main

import (
	"fmt"
	ch5 "go_start/src/countwordsandimages/v1"
)

func main() {
	// url := "http://caibaojian.com/fedbook/tools/http.html"
	url := "www.baidu.com"
	words, images, _ := ch5.CountWordsAndImages(url)

	fmt.Printf("words: %d\nimages: %d\n", words, images)
}
