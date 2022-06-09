/*
 * File: \grpc\client\main.go                                                  *
 * Project: test                                                               *
 * Created At: Thursday, 2022/06/9 , 00:49:24                                  *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Thursday, 2022/06/9 , 00:50:32                               *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package main

import (
	"context"
	"fmt"
	"log"
	"test/grpc"

	grpc "google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := hello.NewHelloServiceClient(conn)
	reply, err := client.Hello(context.Background(), &hello.String{Value: "hello"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply.GetValue())
}
