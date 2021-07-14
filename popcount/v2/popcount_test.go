/*
 * @Descripttion:
 * @version: v0.1
 * @Author: Elon C
 * @Date: 2021-07-14 20:06:40
 * @LastEditors: Elon C
 * @LastEditTime: 2021-07-14 20:39:31
 * @FilePath: \Golang_Starting\popcount\v2\popcount_test.go
 */
package popcount

import "testing"

func TestPopCount(t *testing.T) {
	tests := []struct {
		num uint64
		pop int
	}{
		{0b11, 2},
		{0b10, 1},
	}
	for _, test := range tests {
		if PopCount(test.num) != test.pop {
			t.Errorf("PopCount(%d) = %d\n", test.num, test.pop)
		}
	}
}

func BenchmarkPopCount(b *testing.B) {
	var i uint64
	for i = 0; i < uint64(b.N); i++ {
		PopCount(i)
	}
}
