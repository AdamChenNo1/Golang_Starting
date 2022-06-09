/*
 * File: \go_concurrency\ch2\pool\non-pool-network-daemon\main.go              *
 * Project: Golang_Starting                                                    *
 * Created At: Thursday, 2022/05/26 , 17:37:34                                 *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Thursday, 2022/05/26 , 17:37:48                              *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package main

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

func connectToService() any {
	time.Sleep(1 * time.Second)
	return struct{}{}
}

func startNetworkDaemon() *sync.WaitGroup {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		server, err := net.Listen("tcp", "localhost:8080")
		if err != nil {
			log.Fatalf("cannot listen: %v", err)
		}
		defer server.Close()

		wg.Done()

		for {
			conn, err := server.Accept()
			if err != nil {
				log.Printf("cannot accept connection: %v", err)
				continue
			}

			connectToService()
			fmt.Fprintln(conn, " ")
			conn.Close()
		}
	}()

	return &wg
}