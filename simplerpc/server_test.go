/*
 * File: \main\main.go                                                         *
 * Project: simplerpc                                                          *
 * Created At: Tuesday, 2022/06/14 , 14:39:56                                  *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Tuesday, 2022/06/14 , 19:59:05                               *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package simplerpc

import (
	"encoding/json"
	"fmt"
	"go_start/simplerpc/codec"
	"net"
	"time"
)

func startServer(addr chan string) {
	l, err := net.Listen("tcp", ":0")

	if err != nil {
		panic(fmt.Sprintln("network error:", err))
	}

	fmt.Println("start rpc server on", l.Addr())

	addr <- l.Addr().String()
	Accept(l)
}

func ExampleCodec() {
	addr := make(chan string)
	go startServer(addr)

	// in fact, following code is like a simple rpc client
	conn, _ := net.Dial("tcp", <-addr)
	defer func() { conn.Close() }()

	time.Sleep(time.Second)
	// send options
	json.NewEncoder(conn).Encode(DefaultOption)
	cc := codec.NewGobCodec(conn)
	// send request & receive response
	for i := 0; i < 5; i++ {
		h := &codec.Header{
			ServiceMethod: "Foo.Sum",
			Seq:           uint64(i),
		}
		cc.Write(h, fmt.Sprintf("simplerpc req %d", h.Seq))
		cc.ReadHeader(h)
		var reply string
		cc.ReadBody(&reply)
		fmt.Println("reply:", reply)
	}
	// Output:
	// &{Foo.Sum 1 } simplerpc req 1
	// reply: simplerpc resp 1
	// &{Foo.Sum 2 } simplerpc req 2
	// reply: simplerpc resp 2
	// &{Foo.Sum 3 } simplerpc req 3
	// reply: simplerpc resp 3
	// &{Foo.Sum 4 } simplerpc req 4
	// reply: simplerpc resp 4
}
