/*
 * File: /redis/token_bucket/main.go                                           *
 * Project: go-demo                                                            *
 * Created At: Saturday, 2022/06/11 , 11:31:50                                 *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Saturday, 2022/06/11 , 13:58:19                              *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package ratelimiter

import (
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
)

func ExampleNewScript() {
	fmt.Println(tokenBucketRedisLuaIsLimitedScript.Hash())
	// Output:
	// 1
}

func ExampleDo() {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		// handle error
		log.Fatal(err)
	}
	defer c.Close()

	script1 := redis.NewScript(1, `
		return redis.pcall('set',KEYS[1],ARGV[1])
	`)

	reply, err := script1.Do(c, "tom", 1)

	if err != nil {
		// handle error
		log.Fatal(err)
	}

	// fmt.Printf("%T\n", reply)
	fmt.Println(reply)

	script2 := redis.NewScript(1, `
		return redis.pcall('get',KEYS[1])
	`)

	replies, err := redis.Uint64(script2.Do(c, "tom"))

	if err != nil {
		// handle error
		log.Fatal(err)
	}

	// fmt.Printf("%T\n", reply)
	fmt.Println(replies)

	// Output:
	// OK
	// 1
}
