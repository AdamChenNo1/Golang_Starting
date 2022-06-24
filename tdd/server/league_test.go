/*
 * File: /server/league_test.go                                                *
 * Project: tdd                                                                *
 * Created At: Friday, 2022/06/24 , 07:11:12                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Friday, 2022/06/24 , 10:14:50                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	jsonContentType = "application/json"
)

func TestLeague(t *testing.T) {
	store := stubPlayerStore{}
	s := NewPlayerServer(&store)

	t.Run("it returns 200 on /league", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/league", nil)
		response := httptest.NewRecorder()

		s.ServeHTTP(response, request)
		var got []Player

		err := json.NewDecoder(response.Body).Decode(&got)
		if err != nil {
			t.Fatalf("Unable to parse response from server '%s' into slice of Player, '%v'", response.Body, err)
		}

		AssertResponseStatus(t, response.Code, http.StatusOK)

	})

	t.Run("it returns the league table as JSON", func(t *testing.T) {
		wantedLeague := []Player{
			{"Cler", 32},
			{"Chris", 32},
			{"Tiest", 32},
		}
		store := stubPlayerStore{nil, nil, wantedLeague}
		s := NewPlayerServer(&store)

		request := NewLeagueRequest()
		response := httptest.NewRecorder()

		s.ServeHTTP(response, request)

		got := GetLeagueFromResponse(t, response.Body)

		AssertResponseStatus(t, response.Code, http.StatusOK)
		AssertContentType(t, response, jsonContentType)

		AssertLeague(t, got, wantedLeague)

	})
}
