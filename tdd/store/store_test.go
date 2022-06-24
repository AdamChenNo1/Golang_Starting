/*
 * File: /main_test.go                                                         *
 * Project: tdd                                                                *
 * Created At: Friday, 2022/06/24 , 05:03:43                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Friday, 2022/06/24 , 13:20:59                                *
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
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetrievingThemWithInMemoryStore(t *testing.T) {
	store := NewInMemoryPlayerStore()
	s := server.NewPlayerServer(store)
	player := "Pepper"

	s.ServeHTTP(httptest.NewRecorder(), server.NewPostWinRequest(player))
	s.ServeHTTP(httptest.NewRecorder(), server.NewPostWinRequest(player))
	s.ServeHTTP(httptest.NewRecorder(), server.NewPostWinRequest(player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		s.ServeHTTP(response, server.NewGetScoreRequest(player))
		server.AssertResponseStatus(t, response.Code, http.StatusOK)
		server.AssertResponseBody(t, response.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		s.ServeHTTP(response, server.NewLeagueRequest())
		server.AssertResponseStatus(t, response.Code, http.StatusOK)

		got := server.GetLeagueFromResponse(t, response.Body)
		want := model.League{
			{"Pepper", 3},
		}
		server.AssertLeague(t, got, want)
	})
}

func TestRecordingWinsAndRetrievingThemWithFileSystemStore(t *testing.T) {
	database, cleanDatabase := createTempFile(t, `[]`)
	defer cleanDatabase()
	fstore, err := NewFileSystemPlayerStore(database)
	test.AssertNoError(t, err)

	s := server.NewPlayerServer(fstore)

	player := "Pepper"

	s.ServeHTTP(httptest.NewRecorder(), server.NewPostWinRequest(player))
	s.ServeHTTP(httptest.NewRecorder(), server.NewPostWinRequest(player))
	s.ServeHTTP(httptest.NewRecorder(), server.NewPostWinRequest(player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		s.ServeHTTP(response, server.NewGetScoreRequest(player))
		server.AssertResponseStatus(t, response.Code, http.StatusOK)
		server.AssertResponseBody(t, response.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		s.ServeHTTP(response, server.NewLeagueRequest())
		server.AssertResponseStatus(t, response.Code, http.StatusOK)

		got := server.GetLeagueFromResponse(t, response.Body)
		want := model.League{
			{"Pepper", 3},
		}
		server.AssertLeague(t, got, want)
	})
}
