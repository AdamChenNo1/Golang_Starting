/*
 * File: \src\ch8\bank\main.go                                                 *
 * Project: Golang_Starting                                                    *
 * Created At: Thursday, 2022/05/19 , 23:51:37                                 *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Friday, 2022/05/20 , 12:37:25                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package main

import "sync"

var (
	balance int
	mu      sync.Mutex
	rwmu    sync.RWMutex
)

func Deposit(amount int) {
	mu.Lock()
	defer mu.Unlock()

	deposit(amount)
}

func Balance() int {
	rwmu.Lock()
	defer rwmu.Unlock()
	return balance
}

func deposit(amount int) {
	balance += amount
}

func Withdraw(amount int) bool {
	mu.Lock()
	defer mu.Unlock()

	deposit(-amount)

	if balance < 0 {
		deposit(amount)
		return false
	}

	return true
}
