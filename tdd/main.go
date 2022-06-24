/*
 * File: /http/main.go                                                         *
 * Project: tdd                                                                *
 * Created At: Friday, 2022/06/24 , 03:05:21                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Friday, 2022/06/24 , 06:27:46                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package main

import (
	"go_start/tdd/server"
	"log"
	"net/http"
)

func main() {
	s := &server.PlayerServer{NewInMemoryPlayerStore()}

	if err := http.ListenAndServe(":5000", s); err != nil {
		log.Fatalf("could not listen on port 5000: %v", err)
	}
}
