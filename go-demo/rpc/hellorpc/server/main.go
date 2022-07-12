/*
 * File: \rpc\hellorpc\main.go\main.go                                         *
 * Project: go-demo                                                            *
 * Created At: Wednesday, 2022/06/8 , 22:49:30                                 *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Sunday, 2022/06/26 , 12:34:54                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package main

import (
	hellorpc "go_start/go-demo/rpc"
	"log"
	"net"
	"net/rpc"
)

func main() {
	rpc.RegisterName("HelloService", new(hellorpc.HelloService))

	listen, err := net.Listen("tcp", ":1234")

	if err != nil {
		log.Fatal("ListenTCP error: ", err)
	}

	conn, err := listen.Accept()

	if err != nil {
		log.Fatal("Accept error: ", err)
	}

	rpc.ServeConn(conn)
}
