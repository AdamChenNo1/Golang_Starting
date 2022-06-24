/*
 * File: /server/league_test.go                                                *
 * Project: tdd                                                                *
 * Created At: Friday, 2022/06/24 , 07:11:12                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Friday, 2022/06/24 , 14:10:34                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package server

import (
	"encoding/json"
	"go_start/tdd/model"
	"go_start/tdd/test"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	jsonContentType = "application/json"
)

func TestLeague(t *testing.T) {
	store := test.StubPlayerStore{}
	s := NewPlayerServer(&store)

	t.Run("it returns 200 on /league", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/league", nil)
		response := httptest.NewRecorder()

		s.ServeHTTP(response, request)
		var got model.League

		err := json.NewDecoder(response.Body).Decode(&got)
		if err != nil {
			t.Fatalf("Unable to parse response from server '%s' into slice of Player, '%v'", response.Body, err)
		}

		test.AssertResponseStatus(t, response.Code, http.StatusOK)

	})

	t.Run("it returns the league table as JSON", func(t *testing.T) {
		wantedLeague := model.League{
			{"Cler", 32},
			{"Chris", 32},
			{"Tiest", 32},
		}
		store := test.StubPlayerStore{nil, nil, wantedLeague}
		s := NewPlayerServer(&store)

		request := NewLeagueRequest()
		response := httptest.NewRecorder()

		s.ServeHTTP(response, request)

		got := test.GetLeagueFromResponse(t, response.Body)

		test.AssertResponseStatus(t, response.Code, http.StatusOK)
		test.AssertContentType(t, response, jsonContentType)

		test.AssertLeague(t, got, wantedLeague)

	})
}
