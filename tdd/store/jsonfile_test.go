/*
 * File: /store/jsonfile_test.go                                               *
 * Project: tdd                                                                *
 * Created At: Friday, 2022/06/24 , 11:12:52                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Friday, 2022/06/24 , 13:21:10                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package store

import (
	"go_start/tdd/model"
	"go_start/tdd/server"
	"go_start/tdd/test"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("league from a reader", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t,
			`[
				{"Name":"Cleo","Wins":10},
				{"Name":"Chris","Wins":33}
			]`,
		)

		defer cleanDatabase()

		fstore, err := NewFileSystemPlayerStore(database)
		test.AssertNoError(t, err)

		got := fstore.GetLeague()

		want := model.League{
			{"Cleo", 10},
			{"Chris", 33},
		}

		server.AssertLeague(t, got, want)

		// read again
		got = fstore.GetLeague()
		server.AssertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t,
			`[
				{"Name":"Cleo","Wins":10},
				{"Name":"Chris","Wins":33}
			]`,
		)

		defer cleanDatabase()

		fstore, err := NewFileSystemPlayerStore(database)
		test.AssertNoError(t, err)

		got := fstore.GetPlayerScore("Chris")

		want := 33

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
		AssertEqual(t, got, want)
	})

	t.Run("record wins for existing players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t,
			`[
				{"Name":"Cleo","Wins":10},
				{"Name":"Chris","Wins":33}
			]`,
		)

		defer cleanDatabase()

		fstore, err := NewFileSystemPlayerStore(database)
		test.AssertNoError(t, err)

		fstore.RecordWin("Chris")

		got := fstore.GetPlayerScore("Chris")
		want := 34
		AssertEqual(t, got, want)
	})

	t.Run("record wins for new players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t,
			`[
				{"Name":"Cleo","Wins":10},
				{"Name":"Chris","Wins":33}
			]`,
		)

		defer cleanDatabase()

		fstore, err := NewFileSystemPlayerStore(database)
		test.AssertNoError(t, err)

		fstore.RecordWin("Pepper")

		got := fstore.GetPlayerScore("Pepper")
		want := 1
		AssertEqual(t, got, want)
	})

	t.Run("works with an empty file", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, "")
		defer cleanDatabase()

		_, err := NewFileSystemPlayerStore(database)

		test.AssertNoError(t, err)
	})
}
