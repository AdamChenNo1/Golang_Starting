/*
 * File: \src\ftp\server\main.go                                               *
 * Project: Golang_Starting                                                    *
 * Created At: Wednesday, 2022/05/18 , 23:53:53                                *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Thursday, 2022/05/19 , 00:34:36                              *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

const port = 55

func main() {
	listener, err := net.Listen("tcp", "localhost:"+fmt.Sprintf("%s", port))

	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	// currentWorkingDir := path.Dir("/")
	reader := bufio.NewScanner(conn)
	for reader.Scan() {
		switch reader.Text() {
		case "ls":

		}
	}

}
