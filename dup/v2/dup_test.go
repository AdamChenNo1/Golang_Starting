/*
 * @Descripttion:
 * @version: v0.1
 * @Author: Elon C
 * @Date: 2021-05-22 11:28:47
 * @LastEditors: Elon C
 * @LastEditTime: 2021-05-24 21:31:54
 * @FilePath: \GoPath\src\dup\dup_test.go
 */
package main

import (
	"strings"
	"testing"
)

func TestDup(t *testing.T) {
	var tests = []struct {
		s    string
		want string
	}{}
	// println(tests)
	for _, test := range tests {
		in = strings.NewReader(test.s)
		dup()
	}
}

func ExampleDup(t *testing.T) {
	tests := [...]string{
		"",
		"",
		"1",
		"1",
		"2",
		"a",
		"b",
		"+-/*",
		"+-/*",
	}
	for _, test := range tests {
		in = strings.NewReader(test)
		dup()
	}
	//输出：
	//2	""
	//2	"1"
	//2	"+-/*"
}

func TestMain(m *testing.M) {
	m.Run()
}
