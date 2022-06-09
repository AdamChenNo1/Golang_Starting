/*
 * File: \go_concurrency\ch3\generator\main.go                                 *
 * Project: Golang_Starting                                                    *
 * Created At: Thursday, 2022/05/26 , 20:53:41                                 *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Friday, 2022/05/27 , 18:30:53                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package main

import (
	"fmt"
	pipline "go_start/go_concurrency/ch3"
	"math/rand"
)

func main() {
	repeator()
	// function_repeator()
	// string_repeator()
}

func repeator() {
	done := make(chan any)

	defer close(done)
	for num := range pipline.Take(done, pipline.Repeat(done, 1), 10) {
		fmt.Printf("%d ", num)
	}

}

func function_repeator() {

	done := make(chan any)
	defer close(done)

	rand := func() any {
		return rand.Int()
	}

	for num := range pipline.Take(done, pipline.RepeatFn(done, rand), 10) {
		fmt.Println(num)
	}
}

func string_repeator() {
	done := make(chan any)
	defer close(done)

	var message string
	for token := range pipline.ToString(done, pipline.Take(done, pipline.Repeat(done, "I", "am."), 5)) {
		message += token
	}

	fmt.Printf("message: %s...", message)
}
