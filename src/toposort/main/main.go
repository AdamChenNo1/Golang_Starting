package main

import (
	"fmt"
	ch5 "go_start/src/toposort"
)

// 计算机科学课程及其先决课
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

// 打印计算机科学课程的拓扑排序
func main() {
	for i, course := range ch5.TopoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}
