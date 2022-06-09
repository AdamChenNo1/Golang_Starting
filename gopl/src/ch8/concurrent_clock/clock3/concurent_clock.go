/*
 * File: \src\concurrent_clock\concurent_clock.go                              *
 * Project: Golang_Starting                                                    *
 * Created At: Wednesday, 2022/05/18 , 00:24:38                                *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Wednesday, 2022/05/18 , 21:46:27                             *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package main

import (
	"flag"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	var port string
	flag.StringVar(&port, "port", "8000", "usage: -port=8000 etc")
	flag.Parse()
	listener, err := net.Listen("tcp", "localhost:"+port)

	if err != nil {
		log.Fatal(err)
	}
	log.Printf("listening at %s...\n", port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}

		go handleConn(conn) //并发处理多个连接
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
