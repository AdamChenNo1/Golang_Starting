/*
 * File: \src\ch8\bank\main.go                                                 *
 * Project: Golang_Starting                                                    *
 * Created At: Thursday, 2022/05/19 , 23:51:37                                 *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Thursday, 2022/05/19 , 23:56:25                              *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package main

var (
	deposits = make(chan int) //sending deposit amount
	balances = make(chan int) //receiving balance amount
)

func init() {
	go teller()
}

func Deposit(amount int) {
	deposits <- amount
}

func Balance() int {
	return <-balances
}

func teller() {
	var balance int //balance is limited in teller
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}
