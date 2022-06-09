/*
 * File: \redis\connection\main.go                                             *
 * Project: go-demo                                                            *
 * Created At: Thursday, 2022/06/9 , 22:12:02                                  *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Thursday, 2022/06/9 , 22:15:05                               *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package main

import (
	"github.com/gomodule/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		// handle error
	}
	defer c.Close()
	n, err := conn.Do("APPEND", "key", "value")

}
