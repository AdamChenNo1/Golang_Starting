/*
 * File: \go_concurrency\ch3\pipline\main.go                                   *
 * Project: Golang_Starting                                                    *
 * Created At: Thursday, 2022/05/26 , 20:48:52                                 *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Thursday, 2022/05/26 , 20:52:45                              *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package main

import "fmt"

func main() {
	test1()
}

func test1() {

	multiply := func(values []int, multiplier int) []int {
		multipliedValues := make([]int, len(values))
		for i, v := range values {
			multipliedValues[i] = v * multiplier
		}
		return multipliedValues
	}

	add := func(values []int, additive int) []int {
		addedValues := make([]int, len(values))
		for i, v := range values {
			addedValues[i] = v + additive
		}
		return addedValues
	}

	ints := []int{1, 2, 3, 4}
	for _, v := range add(multiply(ints, 2), 1) {
		fmt.Println(v)
	}
}
