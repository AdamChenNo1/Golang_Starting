/*
 * File: \src\ch8\bank\main.go                                                 *
 * Project: Golang_Starting                                                    *
 * Created At: Thursday, 2022/05/19 , 23:51:37                                 *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Friday, 2022/05/20 , 00:11:27                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package main

type withdraw struct {
	amount int
	result chan<- bool
}

var (
	deposits  = make(chan int) //sending deposit amount
	balances  = make(chan int) //receiving balance amount
	withdraws = make(chan withdraw)
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

func Withdraw(amount int) bool {
	res := make(chan bool)
	withdraws <- withdraw{amount, res}
	if result, ok := <-res; ok {
		return result
	}
	return false
}

func teller() {
	var balance int //balance is limited in teller
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case wit := <-withdraws:
			if wit.amount > balance {
				wit.result <- false
			} else {
				balance -= wit.amount
				wit.result <- true
			}
		}
	}
}
