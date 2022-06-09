package main

import (
	"fmt"
	ch5 "go_start/src/wordfreq"
	"os"
	"sort"
)

func ExampleWordFreq() {
	filename := "C:\\Users\\Elon\\Desktop\\golang\\Golang_Starting\\src\\nodecount\\main.go"

	file, _ := os.Open(filename)
	wordfreqs := ch5.WordFreq(file)
	words := make([]string, 0, len(wordfreqs))

	for word, _ := range wordfreqs {
		words = append(words, word)
	}
	sort.Strings(words)

	for _, word := range words {
		fmt.Printf("%s : %d\n", word, wordfreqs[word])

	}

	//	Output:
	// != : 1
	// "fmt" : 1
	// "go_start/src/nodecount/lib" : 1
	// "golang.org/x/net/html" : 1
	// "os" : 1
	// "outline：%v\n", : 1
	// %d\n", : 1
	// ( : 1
	// ) : 1
	// // : 1
	// := : 2
	// = : 1
	// ch5 : 1
	// ch5.NodeCount(0, : 1
	// count : 1
	// count) : 1
	// doc) : 1
	// doc, : 1
	// err : 2
	// err) : 1
	// fmt.Fprintf(os.Stderr, : 1
	// fmt.Printf("Count : 1
	// func : 1
	// html.Parse(os.Stdin) : 1
	// if : 1
	// import : 1
	// main : 1
	// main() : 1
	// nil : 1
	// nodecount : 1
	// nodes : 1
	// of : 1
	// os.Exit(1) : 1
	// package : 1
	// { : 2
	// } : 2
	// 统计HTML文档树内所有的元素个数 : 1
}
