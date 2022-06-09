/*
 * File: \exercise\ch9\9.1\bank_test.go                                        *
 * Project: Golang_Starting                                                    *
 * Created At: Friday, 2022/05/20 , 00:11:47                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Friday, 2022/05/20 , 00:22:33                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package main

import "testing"

func TestDeposit(t *testing.T) {
	if temp := Balance(); temp != 0 {
		t.Errorf("initial balance should be 0,but got %d\n", temp)
	}
	x := 1
	Deposit(x)
	t.Logf("deposited %d\n", x)
	if temp := Balance(); temp != x {
		t.Errorf("balance should be %d,but got %d\n", x, temp)
	}
}

func TestWithdraw(t *testing.T) {
	if temp := Balance(); temp != 0 {
		t.Errorf("initial balance should be 0,but got %d\n", temp)
	}
	x := 1
	Deposit(x)
	t.Logf("deposited %d\n", x)
	if temp := Balance(); temp != x {
		t.Errorf("balance should be %d,but got %d\n", x, temp)
	}

	y := 1000
	Deposit(y)
	t.Logf("deposited %d\n", y)
	if temp := Balance(); temp != x+y {
		t.Errorf("balance should be %d,but got %d\n", x+y, temp)
	}

	ok := Withdraw(y)
	t.Logf("withdrawed %d\n", y)
	if !ok {
		t.Error("withdraw should succeed,but failed")
	}

	if temp := Balance(); temp != x {
		t.Errorf("balance should be %d,but got %d\n", x, temp)
	}

	if ok := Withdraw(y); ok {
		t.Error("withdraw should fail,but succeeded")
	}
}
