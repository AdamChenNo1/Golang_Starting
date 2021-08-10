/*
 * @Descripttion:
 * @version: v0.1
 * @Author: Elon C
 * @Date: 2021-07-15 18:07:40
 * @LastEditors: Elon C
 * @LastEditTime: 2021-07-16 15:44:03
 * @FilePath: \Golang_Starting\github\main.go
 */
package main

import (
	"fmt"
	github "gostart/github/issues"
	"log"
	"os"
	"strings"
)

func main() {
	var function = os.Args[1]
	if function == "search" {
		var obj = os.Args[2]
		if obj == "issues" {
			vars := os.Args[3:]
			var qStart, sortStart = 0, len(vars) - 1
			for index, item := range vars {
				if strings.HasPrefix(item, "-q=") {
					qStart = index
				}
				if strings.HasPrefix(item, "-s=") {
					sortStart = index
					break
				}
			}
			q := vars[qStart:sortStart]
			result, err := github.SearchIssues(q)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%d issues:\n", result.TotalCount)
			for _, item := range result.Items {
				fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
			}
		}
	}

}
