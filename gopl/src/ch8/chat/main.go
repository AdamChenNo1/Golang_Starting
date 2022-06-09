/*
 * File: \src\ch8\chat\main.go                                                 *
 * Project: Golang_Starting                                                    *
 * Created At: Thursday, 2022/05/19 , 23:21:26                                 *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Thursday, 2022/05/19 , 23:34:10                              *
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

func main() {
	l, err := net.Listen("tcp", "localhost:8000")

	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()

	for {
		conn, err := l.Accept()

		if err != nil {
			log.Print(err)
			continue
		}

		go handleConn(conn)
	}
}

type client chan<- string

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func broadcaster() {
	clients := make(map[client]bool)

	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			clients[cli] = false
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string)

	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}

	leaving <- ch
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
