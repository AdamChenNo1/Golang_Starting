/*
 * File: \go_concurrency\ch1\deadlock\main.go                                  *
 * Project: Golang_Starting                                                    *
 * Created At: Wednesday, 2022/05/25 , 20:09:24                                *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Wednesday, 2022/05/25 , 20:14:59                             *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	type value struct {
		mu    sync.Mutex
		value int
	}

	var wg sync.WaitGroup

	printSum := func(v1, v2 *value) {
		defer wg.Done()
		v1.mu.Lock()         //获取锁，进入临界区
		defer v1.mu.Unlock() //释放锁，退出临界区

		time.Sleep(2 * time.Second) //休眠，模拟工作

		v2.mu.Lock()         //获取锁，进入临界区
		defer v1.mu.Unlock() //释放锁，退出临界区

		fmt.Printf("sum=%v\n", v1.value+v2.value)
	}

	var a, b value

	wg.Add(2)

	go printSum(&a, &b)
	go printSum(&b, &a)

	wg.Wait()
}
