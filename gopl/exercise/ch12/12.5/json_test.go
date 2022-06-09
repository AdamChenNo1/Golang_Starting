/*
 * File: \exercise\ch12\12.2\sexpr_test.go                                     *
 * Project: Golang_Starting                                                    *
 * Created At: Monday, 2022/05/23 , 21:18:21                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Tuesday, 2022/05/24 , 13:04:15                               *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package json

import (
	"bytes"
	"encoding/json"
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

	// var x st = st{5}
	// var y st = st{6}

	tests := []struct {
		v   any
		res string
	}{
		{1, "1"},
		{1.5, "1.5"},
		{true, "true"},
		{false, "false"},
		// {1 + 2i, "1+2i"},
		// {2i, "0+2i"},
		{"it is my fault", `"it is my fault"`},
		{[2]int{1, 2}, "[1, 2]"},
		{[]int{1, 2}, "[1, 2]"},
		{
			struct {
				a string
				b int
			}{
				a: "1",
				b: 2,
			},
			`{"a": "1", "b": 2}`,
		},
		// {map[string]int{"a": 1, "b": 2}, `{"a": 1, "b": 2}`},
		{
			map[string]struct {
				a string
				b int
				c float32
				d []bool
				e struct {
					x string
					y int
				}
			}{
				"first":  {},
				"second": {},
				"third":  {},
				"forth":  {},
			},
			`{}`,
		},
	}
	// {
	// 	map[int]struct{}{1: {}, 2: {}},
	// 	`{"1":{}, "2":{}}`},
	// }
	// {x, "((x 5))"},
	// {y, "((x 6))"},

	for _, test := range tests {
		resb, err := Marshal(test.v)
		if err != nil {
			t.Error(err)
		}
		res := string(resb)

		shouldb, err := json.Marshal(test.v)
		if err != nil {
			t.Error(err)
		}

		should := string(shouldb)

		if res != should {
			t.Errorf("Marshal(%v) = %s,want %s", test.v, res, should)
		}
	}
}

func TestEncode(t *testing.T) {

	// var x i1 = st{5}
	// var y i1 = st{6}

	tests := []struct {
		v   reflect.Value
		res string
	}{
		{reflect.ValueOf(1), "1"},
		{reflect.ValueOf(1.5), "1.5"},
		{reflect.ValueOf(true), "true"},
		{reflect.ValueOf(false), "false"},
		{reflect.ValueOf(1 + 2i), "1+2i"},
		{reflect.ValueOf(2i), "0+2i"},
		{reflect.ValueOf("it is my fault"), "it is my fault"},
		{reflect.ValueOf([2]int{1, 2}), "[1, 2]"},
		{reflect.ValueOf([]int{1, 2}), "[1, 2]"},
		{
			reflect.ValueOf(struct {
				a string
				b int
			}{
				a: "1",
				b: 2,
			}),
			`{"a":1, "b":2}`,
		},
		{reflect.ValueOf(map[string]int{"a": 1, "b": 2}), `{"a": 1, "b": 2}`},
		{reflect.ValueOf(map[int]struct{}{1: {}, 2: {}}), `{"1":{}, "2":{}}`},
		// {reflect.ValueOf(&x).Elem(), `"st"((x 5))`},
		// {reflect.ValueOf(&y).Elem(), `"st"((x 6))`},
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
