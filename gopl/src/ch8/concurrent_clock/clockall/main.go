/*
 * File: \src\concurrent_clock\clockall\main.go                                *
 * Project: Golang_Starting                                                    *
 * Created At: Wednesday, 2022/05/18 , 21:48:25                                *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Wednesday, 2022/05/18 , 23:50:37                             *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	// addrs := flag.Args()
	addrs := os.Args[1:]
	fmt.Println(addrs)
	for _, a := range addrs {
		conn, err := net.Dial("tcp", a[strings.Index(a, "=")+1:])
		log.Println(a)
		if err != nil {
			log.Fatal(err)
		}

		defer conn.Close()

		mustCopy(os.Stdout, conn)
	}
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
