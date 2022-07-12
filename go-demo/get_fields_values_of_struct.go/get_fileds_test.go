/*
 * File: /get_fields_values_of_struct.go/get_fileds_test.go                    *
 * Project: go-demo                                                            *
 * Created At: Tuesday, 2022/06/28 , 06:41:49                                  *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Tuesday, 2022/06/28 , 12:23:35                               *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package main

import (
	"reflect"
	"testing"
)

func TestGetFields2(t *testing.T) {
	type N struct {
		Age int
	}
	type S struct {
		N
		ID    int
		Name  string
		State bool
		Nums  [5]int
	}

	got := get_fields2(S{Nums: [5]int{1, 2, 3, 4, 5}})
	want := []any{
		0, 0, "", false, [5]int{1, 2, 3, 4, 5},
	}
	// want := []any{
	// 	0, "", false,
	// }

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got type %T", got)
		for i, v := range got {
			t.Errorf("got[%d] type %T", i, v)
		}
		t.Errorf("got type %T", got)

		t.Errorf("want type %T", want)
		for i, v := range want {
			t.Errorf("want[%d] type %T", i, v)
		}
		t.Errorf("got %v, want %v", got, want)
	}
}
