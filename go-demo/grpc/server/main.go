/*
 * File: \grpc\server\main.go                                                  *
 * Project: test                                                               *
 * Created At: Thursday, 2022/06/9 , 00:37:47                                  *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Thursday, 2022/06/9 , 00:48:58                               *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package main

import (
	"context"
	"log"
	"net"
	"test/grpc"

	grpc "google.golang.org/grpc"
)

type HelloServiceImpl struct {
}

func (p *HelloServiceImpl) Hello(ctx context.Context, args *hello.String) (*hello.String, error) {
	reply := &hello.String{Value: "hello:" + args.GetValue()}
	return reply, nil
}

func main() {
	grpcServer := grpc.NewServer()
	hello.RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))

	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(lis)
}
