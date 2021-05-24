/*
 * @Descripttion:
 * @version: v0.1
 * @Author: Elon C
 * @Date: 2021-05-23 00:51:18
 * @LastEditors: Elon C
 * @LastEditTime: 2021-05-23 00:57:13
 * @FilePath: \GoPath\src\basename\v1\basename.go
 */
// basename 移除路径部分和.后缀
// e.g.,a => a,a.go => a, a/b/c.go => c,a/b.c.go => b.c
package basename

func basename(s string) string {
	// 将最后一个'/'和之前的部分全部舍弃
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	// 保留最后一个'.'之前的所有内容
	for i := len(s) - 1; i > 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}
