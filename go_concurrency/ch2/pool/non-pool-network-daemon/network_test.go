/*
 * File: \go_concurrency\ch2\pool\non-pool-network-daemon\network_test.go      *
 * Project: Golang_Starting                                                    *
 * Created At: Thursday, 2022/05/26 , 17:38:22                                 *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Thursday, 2022/05/26 , 17:38:43                              *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package main

import (
	"io/ioutil"
	"net"
	"testing"
)

func init() {
	daemonStarted := startNetworkDaemon()
	daemonStarted.Wait()
}

func BenchmarkNetworkRequest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		conn, err := net.Dial("tcp", "localhost:8080")
		if err != nil {
			b.Fatalf("cannot dial host: %v", err)
		}

		if _, err := ioutil.ReadAll(conn); err != nil {
			b.Fatalf("cannot read: %v", err)
		}

		conn.Close()
	}
}
