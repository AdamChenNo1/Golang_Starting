/*
 * File: /test/stub.go                                                         *
 * Project: tdd                                                                *
 * Created At: Friday, 2022/06/24 , 14:07:22                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Friday, 2022/06/24 , 14:11:01                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package test

import "go_start/tdd/model"

type StubPlayerStore struct {
	Scores   map[string]int
	WinCalls []string
	League   model.League
}

func (s *StubPlayerStore) GetLeague() model.League {
	return s.League
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	return s.Scores[name]
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.WinCalls = append(s.WinCalls, name)
}
