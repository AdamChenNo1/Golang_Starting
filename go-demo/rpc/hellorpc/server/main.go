/*
 * File: \rpc\hellorpc\main.go\main.go                                         *
 * Project: test                                                               *
 * Created At: Wednesday, 2022/06/8 , 22:49:30                                 *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Thursday, 2022/06/9 , 00:07:15                               *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package main

import (
	"log"
	"net"
	"net/rpc"
	"test/rpc"
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
