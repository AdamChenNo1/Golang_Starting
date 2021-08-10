/*
 * @Descripttion:
 * @version: v0.1
 * @Author: Elon C
 * @Date: 2021-05-20 22:03:45
 * @LastEditors: Elon C
 * @LastEditTime: 2021-05-20 22:16:20
 * @FilePath: \GoPath\src\TestSplit\split_test.go
 */
//TestSplit 使用基于表的输入和期望输出来测试strings.Split函数
package main

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	var tests = []struct {
		s    string
		sep  string
		want int
	}{
		{"a:b:c", ":", 3},
		{"1 2 3", " ", 3},
	}

	for _, test := range tests {
		words := strings.Split(test.s, test.sep)
		if got := len(words); got != test.want {
			t.Errorf("Split(%q,%q) returned %d words,want %d",
				test.s, test.sep, got, test.want)
		}
	}
}
