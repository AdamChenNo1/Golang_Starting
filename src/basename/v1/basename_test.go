/*
 * @Descripttion:
 * @version: v0.1
 * @Author: Elon C
 * @Date: 2021-05-23 00:57:29
 * @LastEditors: Elon C
 * @LastEditTime: 2021-05-24 00:26:49
 * @FilePath: \GoPath\src\basename\v1\basename_test.go
 */
package basename

import (
	"fmt"
	"testing"
)

func TestBasename(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"", ""},
		{"a", "a"},
		{"a.", "a"},
		{"/a.", "a"},
		{"a.go", "a"},
		{"a.b.go", "a.b"},
		{"a/b/c.go", "c"},
		{"a/b.c.go", "b.c"},
	}

	for _, test := range tests {
		if res := basename(test.s); res != test.want {
			t.Errorf("basename(%s) = %s,want %s", test.s, res, test.want)
		}
	}
}

func ExampleBasename() {
	tests := [...]string{
		"",
		"a",
		"a.",
		"/a.",
		"a.go",
		"a.b.go",
		"a/b/c.go",
		"a/b.c.go",
	}

	for _, test := range tests {
		fmt.Println(basename(test))
	}
	//输出：
	//""
	//"a"
	//"a"
	//"a"
	//"a"
	//"a.b"
	//"c"
	//"b.c"
}
