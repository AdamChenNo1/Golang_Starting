package main

import (
	ch5 "go_start/src/nodecount/lib"
	"os"
	"testing"

	"golang.org/x/net/html"
)

// nodecount 统计HTML文档树内所有的元素个数

func TestNodeCount(t *testing.T) {
	tests := []struct {
		filename   string
		nodenumber int
	}{
		{
			"index.html",
			5,
		},
	}
	for _, test := range tests {
		file, err := os.Open(test.filename)
		if err != nil {
			// fmt.Fprintln(os.Stderr, err)
			t.Fatal(err)
		}

		doc, err := html.Parse(file)
		if err != nil {
			t.Fatalf("NodeCount：%v\n", err)
		}
		count := ch5.NodeCount(0, doc)
		if count != test.nodenumber {
			t.Errorf("%s has %d nodes,wants %d\n", test.filename, count, test.nodenumber)
		}
	}
}
