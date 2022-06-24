/*
 * File: /main_test.go                                                         *
 * Project: tdd                                                                *
 * Created At: Friday, 2022/06/24 , 05:03:43                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Friday, 2022/06/24 , 06:22:30                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package main

import (
	"go_start/tdd/server"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := NewInMemoryPlayerStore()
	s := server.PlayerServer{store}
	player := "Pepper"

	s.ServeHTTP(httptest.NewRecorder(), server.NewPostWinRequest(player))
	s.ServeHTTP(httptest.NewRecorder(), server.NewPostWinRequest(player))
	s.ServeHTTP(httptest.NewRecorder(), server.NewPostWinRequest(player))

	response := httptest.NewRecorder()
	s.ServeHTTP(response, server.NewGetScoreRequest(player))
	server.AssertResponseStatus(t, response.Code, http.StatusOK)
	server.AssertResponseBody(t, response.Body.String(), "3")
}
