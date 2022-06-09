/*
 * @Descripttion:
 * @version: v0.1
 * @Author: Elon C
 * @Date: 2021-05-29 15:21:16
 * @LastEditors: Elon C
 * @LastEditTime: 2021-07-14 20:25:59
 * @FilePath: \Golang_Starting\popcount\v1\popcount.go
 */
package popcount

import "sync"

// pc[i] 是i的种群统计
var (
	initPopCountOnce sync.Once
	pc               [256]byte
)

func initPopCount() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// popCount 返回x的种群统计（置位的个数）
func PopCount(x uint64) int {
	initPopCountOnce.Do(initPopCount)
	return int(pc[byte(x>>0*8)] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
