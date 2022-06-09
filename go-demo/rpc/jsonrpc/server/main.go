/*
 * File: \rpc\jsonrpc\server\main.go                                           *
 * Project: test                                                               *
 * Created At: Wednesday, 2022/06/8 , 23:03:13                                 *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Thursday, 2022/06/9 , 00:04:30                               *
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
	"net/rpc/jsonrpc"
	"test/rpc"
)

func main() {
	rpc.RegisterName("HelloService", new(hellorpc.HelloService))

	listen, err := net.Listen("tcp", ":1234")

	if err != nil {
		log.Fatal("ListenTCP error: ", err)
	}

	for {

		conn, err := listen.Accept()

		if err != nil {
			log.Fatal("Accept error: ", err)
		}

		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
