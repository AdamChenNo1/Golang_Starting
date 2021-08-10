/*
 * @Descripttion:
 * @version: v0.1
 * @Author: Elon C
 * @Date: 2021-05-19 17:22:52
 * @LastEditors: Elon C
 * @LastEditTime: 2021-05-21 01:30:05
 * @FilePath: \GoPath\src\word\v2\Palindrome_test.go
 */
package word

import "testing"

func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"kayak", true},
		{"detartrated", true},
		{"A man,a plan,a canal: Panama", true},
		{"Evil I did dwell;lewd did I live.", true},
		{"Able was I ere I saw Elba", true},
		{"été", true},
		{"Et se resservir,ivresse reste.", true},
		{"palindrome.", false}, //非回文
		{"desserts.", false},   //非回文
	}
	for _, test := range tests {
		if got := IsPalindrome(test.input); got != test.want {
			t.Errorf("IsPalindrome(%q) = %v", test.input, got)
		}
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindrome("A man,a plan,a canal: Panama")
	}
}

func benchmark(b *testing.B, size int) {
	for i := 0; i < size; i++ {
		IsPalindrome("A man,a plan,a canal: Panama")
	}
}

func BenchmarkIsPalindrome10(b *testing.B) {
	benchmark(b, 10)
}

func BenchmarkIsPalindrome100(b *testing.B) {
	benchmark(b, 100)
}

func BenchmarkIsPalindrome1000(b *testing.B) {
	benchmark(b, 1000)
}
