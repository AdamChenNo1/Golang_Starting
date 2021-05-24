/*
 * @Descripttion:
 * @version: v0.1
 * @Author: Elon C
 * @Date: 2021-05-19 17:18:38
 * @LastEditors: Elon C
 * @LastEditTime: 2021-05-21 01:13:52
 * @FilePath: \GoPath\src\word\v2\Palindrome.go
 */
//包word提供了文字游戏相关的工具函数
//忽略字母大小写，以及非字母字符
package word

import "unicode"

//IsPalindrome判断一个字符串是否为回文符串
func IsPalindrome(s string) bool {
	letters := make([]rune, 0, len(s))
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}

	for i := range letters {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}
