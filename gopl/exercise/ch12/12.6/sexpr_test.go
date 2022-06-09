/*
 * File: \exercise\ch12\12.2\sexpr_test.go                                     *
 * Project: Golang_Starting                                                    *
 * Created At: Monday, 2022/05/23 , 21:18:21                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Monday, 2022/05/23 , 23:48:11                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package sexpr

import (
	"bytes"
	"reflect"
	"testing"
)

type st struct {
	x int
}

type i1 interface {
	Compare(i1) bool
}

func (s st) Compare(a i1) bool {
	return s.x < a.(st).x
}

func TestMarshal(t *testing.T) {

	var x st = st{5}
	var y st = st{6}

	tests := []struct {
		v   any
		res string
	}{
		{1, "1"},
		{1.5, "1.5"},
		{true, "t"},
		{false, "nil"},
		{1 + 2i, "#(1,2)"},
		{2i, "#(0,2)"},
		{"it is my fault", "it is my fault"},
		{[2]int{1, 2}, "(1 2)"},
		{[]int{1, 2}, "(1 2)"},
		{
			struct {
				a string
				b int
			}{
				a: "1",
				b: 2,
			},
			"((a 1) (b 2))",
		},
		{map[string]int{"a": 1, "b": 2}, "((a 1) (b 2))"},
		{map[int]struct{}{1: {}, 2: {}}, "((1 ()) (2 ()))"},
		{map[int]struct{}{1: {}, 2: {}}, "((1 ()) (2 ()))"},
		{x, "((x 5))"},
		{y, "((x 6))"},
	}

	for _, test := range tests {
		resb, err := Marshal(test.v)
		res := string(resb)
		if err != nil {
			t.Error(err)
		}

		if res != test.res {
			t.Errorf("Marshal(%v) = %s,want %q", test.v, res, test.res)
		}
	}
}

func TestEncode(t *testing.T) {

	var x i1 = st{5}
	var y i1 = st{6}

	tests := []struct {
		v   reflect.Value
		res string
	}{
		{reflect.ValueOf(1), "1"},
		{reflect.ValueOf(1.5), "1.5"},
		{reflect.ValueOf(true), "t"},
		{reflect.ValueOf(false), "nil"},
		{reflect.ValueOf(1 + 2i), "#(1,2)"},
		{reflect.ValueOf(2i), "#(0,2)"},
		{reflect.ValueOf("it is my fault"), "it is my fault"},
		{reflect.ValueOf([2]int{1, 2}), "(1 2)"},
		{reflect.ValueOf([]int{1, 2}), "(1 2)"},
		{
			reflect.ValueOf(struct {
				a string
				b int
			}{
				a: "1",
				b: 2,
			}),
			"((a 1) (b 2))",
		},
		{reflect.ValueOf(map[string]int{"a": 1, "b": 2}), "((a 1) (b 2))"},
		{reflect.ValueOf(map[int]struct{}{1: {}, 2: {}}), "((1 ()) (2 ()))"},
		{reflect.ValueOf(&x).Elem(), `"st"((x 5))`},
		{reflect.ValueOf(&y).Elem(), `"st"((x 6))`},
	}

	for _, test := range tests {
		var buf bytes.Buffer
		err := encode(&buf, test.v)
		if err != nil {
			t.Error(err)
		}

		res := buf.String()

		if res != test.res {
			t.Errorf("Marshal(%v) = %q,want %q", test.v, res, test.res)
		}
	}
}
