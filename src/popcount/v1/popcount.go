/*
 * @Descripttion:
 * @version: v0.1
 * @Author: Elon C
 * @Date: 2021-07-14 20:25:55
 * @LastEditors: Elon C
 * @LastEditTime: 2021-07-14 20:33:11
 * @FilePath: \Golang_Starting\popcount\v2\popcount.go
 */
package popcount

func PopCount(x uint64) int {
	pop := 0
	for i := 0; i < 64; i++ {
		if (x & 1) != 0 {
			pop += 1
		}
		x >>= 1
	}
	return pop
}
