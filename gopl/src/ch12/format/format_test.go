/*
 * File: \src\ch12\format\format_test.go                                       *
 * Project: Golang_Starting                                                    *
 * Created At: Monday, 2022/05/23 , 17:13:28                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Monday, 2022/05/23 , 17:22:37                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package format

import (
	"fmt"
	"testing"
	"time"
)

func ExampleAny() {
	var x int64

	var d time.Duration = 1 * time.Second

	fmt.Println(Any(x))
	fmt.Println(Any(d))
	fmt.Println(Any([]int64{x}))
	fmt.Println(Any([]time.Duration{d}))

	// Output:
	// 0
	// 1000000000
	// []int64 0xc000014310
	// []time.Duration 0xc000014318
}

func TestAny(t *testing.T) {
	tests := []struct {
		a   any
		res string
	}{
		{
			1, "1",
		},
		{
			int8(1), "1",
		},
		{
			int16(1), "1",
		},
		{
			int32(1), "1",
		},
		{
			int64(1), "1",
		},
		{
			uint(1), "1",
		},
		{
			uint8(1), "1",
		},
		{
			uint16(1), "1",
		},
		{
			uint32(1), "1",
		},
		{
			uint64(1), "1",
		},
		{
			float32(1.0000), "1",
		},
		{
			float64(1.0000), "1",
		},
	}

	for _, test := range tests {
		if res := Any(test.a); res != test.res {
			t.Errorf("Any(%v) = %s,want %s\n", test.a, res, test.res)
		}
	}
}
