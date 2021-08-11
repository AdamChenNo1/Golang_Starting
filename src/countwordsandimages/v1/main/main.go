package main

import (
	"fmt"
	ch5 "go_start/src/countwordsandimages/v1"
	"os"
)

func main() {
	// url := "http://caibaojian.com/fedbook/tools/http.html"
	// url := "www.baidu.com"
	url := os.Args[1]

	words, images, err := ch5.CountWordsAndImages(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "countwordsandimages: %v", err)
		os.Exit(1)
	}

	fmt.Printf("words: %d\nimages: %d\n", words, images)
}
