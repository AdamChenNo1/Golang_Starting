/*
 * File: /reflect_value/reflect_test.go                                        *
 * Project: go-demo                                                            *
 * Created At: Tuesday, 2022/06/28 , 09:12:20                                  *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Tuesday, 2022/06/28 , 09:15:13                               *
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

func TestArrayIndex(t *testing.T) {
	a := [5]int{1, 2, 3, 4, 5}

	want := 1
	got := reflect.ValueOf(a).Index(0).Int()

	if got != int64(want) {
		t.Errorf("got %d, want %d", got, want)
	}
}
