/*
 * File: \src\ch8\bank\main.go                                                 *
 * Project: Golang_Starting                                                    *
 * Created At: Thursday, 2022/05/19 , 23:51:37                                 *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Friday, 2022/05/20 , 12:24:24                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package main

var (
	balance int
	sema    = make(chan struct{}, 1)
)

func Deposit(amount int) {
	sema <- struct{}{} //obtain token
	balance += amount
	<-sema //release token
}

func Balance() int {
	sema <- struct{}{} //obtain token
	b := balance
	<-sema //release token
	return b
}
