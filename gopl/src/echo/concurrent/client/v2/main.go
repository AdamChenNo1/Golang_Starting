/*
 * File: \src\echo\concurrent\client\v2\main.go                                *
 * Project: Golang_Starting                                                    *
 * Created At: Thursday, 2022/05/19 , 10:51:06                                 *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Thursday, 2022/05/19 , 23:37:01                              *
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
	"os"
)

func main() {
	// laddr, err := net.ResolveTCPAddr("tcp", "localhost:8010")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	raddr, err := net.ResolveTCPAddr("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.DialTCP("tcp", nil, raddr)
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})

	go func() {
		_, err := io.Copy(os.Stdout, conn)
		log.Fatal(err)
		log.Println("done")

		done <- struct{}{}
	}()

	mustCopy(conn, os.Stdin)

	conn.CloseWrite()
	<-done
	// if ch == struct{}{} {
	// 	conn.CloseRead()
	// }
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
