/*
 * File: \src\concurrent_clock\concurent_clock.go                              *
 * Project: Golang_Starting                                                    *
 * Created At: Wednesday, 2022/05/18 , 00:24:38                                *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Wednesday, 2022/05/18 , 11:48:20                             *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")

	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}

		handleConn(conn) //一次处理一个连接
	}
}

func handleConn(c net.Conn) {
	defer c.Close()

	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))

		if err != nil {
			return
		}

		time.Sleep(1 * time.Second)
	}
}
