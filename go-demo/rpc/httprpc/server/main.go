/*
 * File: \rpc\httprpc\server\main.go                                           *
 * Project: go-demo                                                            *
 * Created At: Thursday, 2022/06/9 , 00:02:57                                  *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Monday, 2022/06/20 , 21:48:04                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package main

import (
	hellorpc "go-demo/rpc"
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {

	rpc.RegisterName("HelloService", new(hellorpc.HelloService))

	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {

		var conn io.ReadWriteCloser = struct {
			io.Writer

			io.ReadCloser
		}{

			ReadCloser: r.Body,

			Writer: w,
		}

		rpc.ServeCodec(jsonrpc.NewServerCodec(conn))

	})

	http.ListenAndServe(":1234", nil)

}
