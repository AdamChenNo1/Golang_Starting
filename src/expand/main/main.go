package main

import (
	"fmt"
	ch2 "go_start/src/expand"
)

func main() {
	s := "$foo"
	fmt.Printf("s = %s\n", s)
	as := ch2.Expand(s, func(s string) string {
		return "$" + s
	})
	fmt.Printf("Expand(s,f) = %s\n", as)
}
