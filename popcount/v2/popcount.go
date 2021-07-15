/*
 * @Descripttion:
 * @version: v0.1
 * @Author: Elon C
 * @Date: 2021-07-14 20:25:55
 * @LastEditors: Elon C
 * @LastEditTime: 2021-07-15 16:51:49
 * @FilePath: \Golang_Starting\popcount\v3\popcount.go
 */
package popcount

func PopCount(x uint64) int {
	pop := 0
	for int(x) != 0 {
		pop += int(x) & 1
		x >>= 1
	}
	return pop
}
