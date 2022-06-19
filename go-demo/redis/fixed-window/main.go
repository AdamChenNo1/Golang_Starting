/*
 * File: \redis\token_bucket.go                                                *
 * Project: go-demo                                                            *
 * Created At: Thursday, 2022/06/9 , 17:24:35                                  *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Thursday, 2022/06/9 , 17:04:29                               *
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
	"math/rand"
	"time"

	"github.com/gomodule/redigo/redis"
)

func IsActionAllowed(userId, actionKey string, period, maxCount int64) bool {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		// handle error
		log.Fatal(err)
	}
	defer c.Close()

	key := fmt.Sprintf("hist:%s%s", userId, actionKey)
	// fmt.Println(key)
	now_ts := time.Now().UnixMilli() // 毫秒时间戳

	c.Send("MULTI")
	c.Send("ZADD", key, now_ts, now_ts) // value和score都使用毫秒时间戳
	// c.Flush()
	// fmt.Println(c.Receive())
	// c.Receive()

	// fmt.Println(now_ts-period*1000)
	c.Send("ZREMRANGEBYSCORE", key, 0, now_ts-period*1000) // 移除时间窗口之前的行为记录，剩下的都是时间窗口内的
	// c.Flush()
	// fmt.Println(c.Receive())
	c.Send("ZCARD", key) // 获取窗口内的行为数量

	c.Send("EXPIRE", key, period+1) // 设置zset过期时间，避免冷用户持续占用内存，过期时间应该等于时间窗口的长度，再多宽限ls

	// c.Flush()
	// current_count, err := c.Receive()

	// c.Flush()
	// fmt.Println(c.Receive())
	replies, err := redis.Values(c.Do("EXEC")) // 批量执行

	if err != nil {
		// handle error
		log.Fatal(err)
	}

	current_count := replies[2].(int64)
	// fmt.Printf("%T: ", current_count)
	fmt.Println(current_count)
	// fmt.Print("-----------------------")
	return current_count <= maxCount
	// return true
}

func main() {
	for i := 0; i < 20; i++ {
		fmt.Println(IsActionAllowed("laoqian", "reply", 2, 5))
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}
