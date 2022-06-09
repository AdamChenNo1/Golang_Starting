package main

import (
	"fmt"
	ch5 "go_start/src/countwordsandimages/v1"
)

func ExampleCountWordsAndImages() {
	url := "www.baidu.com"
	words, images, _ := ch5.CountWordsAndImages(url)
	fmt.Printf("words: %d\nimages: %d\n", words, images)
	// Output:
	// words: 6489
	// images: 15
}
