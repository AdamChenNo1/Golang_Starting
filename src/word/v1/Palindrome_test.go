/*
 * @Descripttion:
 * @version: v0.1
 * @Author: Elon C
 * @Date: 2021-05-19 17:22:52
 * @LastEditors: Elon C
 * @LastEditTime: 2021-05-19 17:34:53
 * @FilePath: \GoPath\src\word\Palindrome_test.go
 */
package word

import "testing"

func TestPalindrome(t *testing.T) {
	if !IsPalindrome("xaalaax") {
		t.Error(`IsPalindrome("xaalaax") = false`)
	}
	if IsPalindrome("askffg") {
		t.Error(`IsPalindrome("askffg") = true`)
	}
}

func TestFrenchPalindrome(t *testing.T) {
	if !IsPalindrome("été") {
		t.Error(`IsPalindrome("été") = false`)
	}
}

func TestCanalPalindrome(t *testing.T) {
	input := "A man,a plan,a canal:	Panama"
	if !IsPalindrome(input) {
		t.Errorf(`IsPalindrome(%q) = false`, input)
	}
}
