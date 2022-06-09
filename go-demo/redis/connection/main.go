/*
 * File: \redis\connection\main.go                                             *
 * Project: go-demo                                                            *
 * Created At: Thursday, 2022/06/9 , 22:12:02                                  *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Thursday, 2022/06/9 , 15:33:34                               *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package main

import (
	"github.com/gomodule/redigo/redis"
	"log"
	"fmt"
)

func main() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		// handle error
		log.Fatal(err)
	}
	defer conn.Close()
	
	n, err := conn.Do("SET", "APPLE", "1")
	if err != nil {
		log.Fatal(err)
		// handle error
	}

	fmt.Println(n)

	n, err = conn.Do("GET", "APPLE")
	if err != nil {
		log.Fatal(err)
		// handle error
	}

	fmt.Printf("%T\n",n)
}
