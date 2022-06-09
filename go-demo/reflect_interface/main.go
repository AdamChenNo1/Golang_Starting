/*
 * File: \test\reflect_interface\main.go                                       *
 * Project: Golang_Starting                                                    *
 * Created At: Monday, 2022/05/23 , 22:14:20                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Monday, 2022/05/23 , 23:14:03                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package main

import (
	"fmt"
	"reflect"
)

func main() {
	st1 := s1{}
	st2 := s1{}

	var x, y i1
	x = st1
	y = st2
	z := s1{}

	fmt.Println(reflect.TypeOf(x))
	fmt.Println(reflect.TypeOf(x).Name())
	fmt.Println(reflect.ValueOf(x).String())
	fmt.Println(reflect.TypeOf(y))

	switch reflect.TypeOf(x).Kind() {
	case reflect.Interface:
		fmt.Println("interface")
	default:
		fmt.Println("non-interface")
	}

	fmt.Println(reflect.ValueOf(&x).Elem().Kind())
	fmt.Println(reflect.ValueOf(&z).Elem().Kind())
	fmt.Println(reflect.ValueOf(&x).Kind())
}

type i1 interface {
	Compare() bool
}

type s1 struct {
}

func (s s1) Compare() bool {
	return true
}
