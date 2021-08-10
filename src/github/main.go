/*
 * @Descripttion:
 * @version: v0.1
 * @Author: Elon C
 * @Date: 2021-07-15 18:07:40
 * @LastEditors: Elon C
 * @LastEditTime: 2021-07-16 16:45:33
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
			var qStart, sortStart = 0, len(vars)
			for index, item := range vars {
				if strings.HasPrefix(item, "-q=") {
					qStart = index
					vars[index] = strings.TrimPrefix(vars[index], "-q=")
				}
				if strings.HasPrefix(item, "-s=") {
					sortStart = index
					vars[index] = strings.TrimPrefix(vars[index], "-s=")
					break
				}
			}
			q := vars[qStart:sortStart]
			sort := vars[sortStart:len(vars)]
			varMap := make(map[string][]string)
			varMap["q"] = q
			varMap["sort"] = sort
			result, err := github.SearchIssues(varMap)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%d issues:\n", result.TotalCount)
			for _, item := range result.Items {
				fmt.Printf("#%-5d %9.9s %.55s %-.32s\n", item.Number, item.User.Login, item.Title, item.CreatedAt)
			}
		}
	}

}
