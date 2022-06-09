package main

import (
	"fmt"
	ch2 "go_start/src/expand"
	"testing"
)

func TestExpand(t *testing.T) {
	tests := []struct {
		s   string
		f   func(string) string
		res string
	}{
		{
			"$foo",
			func(s string) string {
				return "$" + s
			},
			"$foo",
		},
		{
			"$foofoo",
			func(s string) string {
				return "$"
			},
			"$",
		},
		{
			"$foofoofoo",
			func(s string) string {
				return "$"
			},
			"$",
		},
	}

	for _, test := range tests {
		fmt.Println()
		if text := ch2.Expand(test.s, test.f); test.res != text {
			t.Errorf("expand(%s,f) = %s,want %s", test.s, text, test.res)
		}
	}
}
