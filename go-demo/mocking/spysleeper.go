/*
 * File: /mocking/sleeper.go                                                   *
 * Project: go-demo                                                            *
 * Created At: Sunday, 2022/06/26 , 13:09:30                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Monday, 2022/06/27 , 09:16:31                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package main

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type CountDownOperationSpy struct {
	Calls []string
}

func (s *CountDownOperationSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountDownOperationSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const sleep = "sleep"
const write = "write"
