/*
 * File: /http/http_test.go                                                    *
 * Project: tdd                                                                *
 * Created At: Friday, 2022/06/24 , 02:52:01                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Friday, 2022/06/24 , 09:49:48                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type stubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   []Player
}

func (s *stubPlayerStore) GetLeague() []Player {
	return s.league
}

func (s *stubPlayerStore) GetPlayerScore(name string) int {
	return s.scores[name]
}

func (s *stubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func TestGETPlayerServer(t *testing.T) {
	store := &stubPlayerStore{
		scores: map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
	}
	server := NewPlayerServer(store)

	t.Run("returns Pepper's score", func(t *testing.T) {
		request := NewGetScoreRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "20"

		AssertResponseStatus(t, response.Code, http.StatusOK)
		AssertResponseBody(t, got, want)

	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		request := NewGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "10"

		AssertResponseStatus(t, response.Code, http.StatusOK)
		AssertResponseBody(t, got, want)
	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		request := NewGetScoreRequest("Apollo")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		AssertResponseStatus(t, response.Code, http.StatusNotFound)
	})
}

func TestStoreWins(t *testing.T) {
	store := stubPlayerStore{
		scores: map[string]int{},
	}
	server := NewPlayerServer(&store)

	t.Run("it returns accepted on POST", func(t *testing.T) {
		request := NewPostWinRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		AssertResponseStatus(t, response.Code, http.StatusAccepted)
	})

	t.Run("it records wins when POST", func(t *testing.T) {
		player := "Pepper"

		request := NewPostWinRequest("Pepper")

		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		AssertResponseStatus(t, response.Code, http.StatusAccepted)

		if len(store.winCalls) != 2 {
			t.Errorf("got %v calls to RecordWin, want %v", len(store.winCalls), 2)
		}

		if store.winCalls[0] != player {
			t.Errorf("did not store correct winner, got %v, want %v", store.winCalls[0], player)
		}
	})
}
