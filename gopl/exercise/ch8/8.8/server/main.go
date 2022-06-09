/*
 * File: \src\echo\concurrent\server\main.go                                   *
 * Project: Golang_Starting                                                    *
 * Created At: Thursday, 2022/05/19 , 00:47:19                                 *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Thursday, 2022/05/19 , 19:16:01                              *
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
	"strings"
	"time"
)

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("listening on 8000")
	defer l.Close()

	for {
		conn, err := l.Accept()

		if err != nil {
			log.Fatal(err)
		}

		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	tick := time.Tick(1 * time.Second)
	log.Printf("accepted %s from %s\n", c.RemoteAddr().Network(), c.RemoteAddr().String())
	input := bufio.NewScanner(c)
	data := make(chan string)
	go func() {
		for input.Scan() {
			data <- input.Text()
		}
	}()

	for countdown := 10; countdown > 0; {
		select {
		case <-tick:
			countdown--
			if countdown == 0 {
				log.Printf("detached %s from %s\n", c.RemoteAddr().Network(), c.RemoteAddr().String())
				c.Close()
				return
			}
		case s := <-data:
			countdown = 10
			echo(c, s, 1*time.Second)
		}
	}

}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}
