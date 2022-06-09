/*
 * @Descripttion:
 * @version: v0.1
 * @Author: Elon C
 * @Date: 2021-05-19 18:28:54
 * @LastEditors: Elon C
 * @LastEditTime: 2021-05-23 00:28:59
 * @FilePath: \GoPath\src\echo\echo_test.go
 */
package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestEcho(t *testing.T) {
	var tests = []struct {
		newline bool
		sep     string
		args    []string
		want    string
	}{
		{true, "", []string{}, "\n"},
		{false, "", []string{}, ""},
		{true, "\t", []string{"one", "two", "three"}, "one\ttwo\tthree\n"},
		{true, ",", []string{"a", "b", "c"}, "a,b,c\n"},
		{true, ":", []string{"1", "2", "3"}, "1:2:3\n"},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("echo(%v,%q,%q)", test.newline, test.sep, test.args)
		out = new(bytes.Buffer) //捕获的输出
		if err := echo(test.newline, test.sep, test.args); err != nil {
			t.Errorf("%s failed: %v", descr, err)
			continue
		}
		got := out.(*bytes.Buffer).String()
		if got != test.want {
			t.Errorf("%s = %q, want %q", descr, got, test.want)
		}
	}
}

func ExampleEcho() {
	tests := []struct {
		newline bool
		sep     string
		args    []string
	}{
		{true, "", []string{}},
		{false, "", []string{}},
		{true, "\t", []string{"one", "two", "three"}},
		{true, ",", []string{"a", "b", "c"}},
		{true, ":", []string{"1", "2", "3"}},
	}
	for _, test := range tests {
		echo(test.newline, test.sep, test.args)
		fmt.Printf("%v", out.(*bytes.Buffer).String())
	}
	//输出：
	//"\n"
	//""
	//"one\ttwo\tthree\n"
	//"a,b,c\n"
	//"1:2:3\n"
}

func BenchmarkEcho(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo(true, "\t", []string{"one", "two", "three"})
	}
}
