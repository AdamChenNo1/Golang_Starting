/*
 * @Descripttion:
 * @version: v0.1
 * @Author: Elon C
 * @Date: 2021-07-15 18:07:40
 * @LastEditors: Elon C
 * @LastEditTime: 2021-07-15 19:57:28
 * @FilePath: \Golang_Starting\github\main.go
 */
package main

import (
	"fmt"
	github "gostart/github/issues"
	"log"
	"os"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}
