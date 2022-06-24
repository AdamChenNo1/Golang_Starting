/*
 * File: /http/main.go                                                         *
 * Project: tdd                                                                *
 * Created At: Friday, 2022/06/24 , 03:05:21                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Friday, 2022/06/24 , 16:11:50                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package main

import (
	"go_start/tdd/config"
	"go_start/tdd/server"
	"go_start/tdd/store"
	"log"
	"net/http"
)

func main() {
	fstore, err := store.FileSystemPlayerStoreFromFile(config.DbFileName)

	if err != nil {
		log.Fatalf("problem creating file system player store, %v", err)
	}

	s := server.NewPlayerServer(fstore)

	if err := http.ListenAndServe(":5000", s); err != nil {
		log.Fatalf("could not listen on port 5000: %v", err)
	}
}
