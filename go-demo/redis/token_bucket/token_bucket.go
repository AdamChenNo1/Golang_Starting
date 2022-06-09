/*
 * File: \redis\token_bucket.go                                                *
 * Project: test                                                               *
 * Created At: Thursday, 2022/06/9 , 17:24:35                                  *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Thursday, 2022/06/9 , 22:01:08                               *
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

	"github.com/gomodule/redigo/redis"
)

func IsActionAllowed(userId, actionKey string, period, maxCount int) bool {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		// handle error
		log.Fatal(err)
	}
	defer c.Close()
	return true
}

func main() {

	for i := 0; i < 20; i++ {
		fmt.Println(IsActionAllowed("laoqian", "reply", 60, 5))
	}
}
