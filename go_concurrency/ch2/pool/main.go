/*
 * File: \go_concurrency\ch2\pool\main.go                                      *
 * Project: Golang_Starting                                                    *
 * Created At: Thursday, 2022/05/26 , 17:05:57                                 *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Thursday, 2022/05/26 , 17:42:18                              *
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
	"sync"
	"time"
)

func main() {
	// test1()
	test2()
}

func test1() {

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

}

func test2() {
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
	// 4 calculators were created.
}

func warmServiceConnCache() *sync.Pool {
	p := &sync.Pool{
		New: connectToService,
	}

	for i := 0; i < 10; i++ {
		p.Put(p.New())
	}

	return p
}

func connectToService() any {
	time.Sleep(1 * time.Second)
	return struct{}{}
}

func startNetworkDaemon() *sync.WaitGroup {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		connPool := warmServiceConnCache()
		server, err := net.Listen("tcp", "localhost:8080")
		if err != nil {
			log.Fatalf("cannot listen: %v", err)
		}
		defer server.Close()

		wg.Done()

		for {
			conn, err := server.Accept()
			if err != nil {
				log.Printf("cannot accept connection: %v", err)
				continue
			}

			svcConn := connPool.Get()
			fmt.Fprintln(conn, " ")
			connPool.Put(svcConn)
			conn.Close()
		}
	}()

	return &wg
}
