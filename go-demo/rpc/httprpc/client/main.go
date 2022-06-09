/*
 * File: \rpc\hellorpc\client\main.go                                          *
 * Project: test                                                               *
 * Created At: Wednesday, 2022/06/8 , 22:55:06                                 *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Wednesday, 2022/06/8 , 23:07:44                              *
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
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:1234")

	if err != nil {
		log.Fatal("Dialing: ", err)
	}

	var reply string
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	err = client.Call("HelloService.Hello", "hello", &reply)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}
