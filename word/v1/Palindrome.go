/*
 * @Descripttion:
 * @version: v0.1
 * @Author: Elon C
 * @Date: 2021-05-19 17:18:38
 * @LastEditors: Elon C
 * @LastEditTime: 2021-05-19 17:21:44
 * @FilePath: \GoPath\src\word\Palindrome.go
 */
//包word提供了文字游戏相关的工具函数
package word

//IsPalindrome判断一个字符串是否为回文符串
func IsPalindrome(s string) bool {
	for i := range s {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
