/*
 * File: \go_concurrency\ch2\pool\pool_test.go                                 *
 * Project: Golang_Starting                                                    *
 * Created At: Thursday, 2022/05/26 , 17:08:07                                 *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Thursday, 2022/05/26 , 17:42:38                              *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"sync"
	"testing"
)

func Example1() {
	myPool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating new instance.")
			return struct{}{}
		},
	}

	myPool.Get() // 在这里调用Pool的get方法。这些调用将执行Pool中定义的New函数，因为实例还没有实例化。
	instance := myPool.Get()
	myPool.Put(instance) // 将先前检索到的实例放在池中，这就增加了实例的可用数量。
	myPool.Get()         // 在执行此调用时，将重用以前分配的实例并将其放回池中。New将不会被调用。

	// Output
	// Creating new instance.
	// Creating new instance.
}

func Example2() {
	var numCaclcsCreated int
	calcPool := &sync.Pool{
		New: func() any {
			numCaclcsCreated += 1
			mem := make([]byte, 1024)
			return &mem
		},
	}

	// 用4KB初始化pool
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())

	const numWorkers = 1024 * 1024
	var wg sync.WaitGroup
	wg.Add(numWorkers)
	for i := numWorkers; i > 0; i-- {
		go func() {
			defer wg.Done()
			mem := calcPool.Get().(*[]byte)
			defer calcPool.Put(mem)

			// 做一些有趣的假设，但是很快就会用这个内存完成
		}()
	}

	wg.Wait()
	fmt.Printf("%d calculators were created.", numCaclcsCreated)

	// Output
	// 7 calculators were created.
}

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
