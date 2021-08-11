package ch2

import "strings"

func Expand(s string, f func(string) string) string {
	for strings.Contains(s, "$foo") {
		if s == strings.ReplaceAll(s, "$foo", f("foo")) {
			break
		}
		s = strings.ReplaceAll(s, "$foo", f("foo"))
	}
	return s
}
