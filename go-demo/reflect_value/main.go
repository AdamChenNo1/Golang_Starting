/*
 * File: \test\reflect_value\main.go                                           *
 * Project: Golang_Starting                                                    *
 * Created At: Monday, 2022/05/23 , 19:07:11                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Monday, 2022/05/23 , 19:31:17                                *
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
	type ss struct {
		jack string
	}
	a := ss{jack: "hello"}
	value := reflect.ValueOf(a)
	fmt.Println(value.Kind().String())
	fmt.Println(value.Type())
	fmt.Println(value.Type().Name())
	fmt.Println(value.Type().String())
	for i, l := 0, value.NumField(); i < l; i++ {
		fmt.Printf("%s:%s", value.Type().Field(i).Name, value.Field(i))
	}
}
