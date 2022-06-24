/*
 * File: /cmd/cli/cli_test.go                                                  *
 * Project: tdd                                                                *
 * Created At: Friday, 2022/06/24 , 14:05:53                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Friday, 2022/06/24 , 15:53:40                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package main

import (
	"go_start/tdd/test"
	"strings"
	"testing"
)

func TestCli(t *testing.T) {
	t.Run("record chris win from user input", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playStore := test.StubPlayerStore{}

		cli := NewCLI(&playStore, in)
		cli.PlayPoker()

		AssertPlayerWin(t, &playStore, "Chris")
	})

	t.Run("record Cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")
		playStore := test.StubPlayerStore{}

		cli := NewCLI(&playStore, in)
		cli.PlayPoker()

		AssertPlayerWin(t, &playStore, "Cleo")
	})
}

func AssertPlayerWin(t *testing.T, store *test.StubPlayerStore, winner string) {
	t.Helper()

	if len(store.WinCalls) != 1 {
		t.Fatalf("got %d calls to RecordWin, want %d", len(store.WinCalls), 1)
	}

	if store.WinCalls[0] != winner {
		t.Errorf("didn't store correct winner, got %s,want %s", store.WinCalls[0], winner)
	}
}
