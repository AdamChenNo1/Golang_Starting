/*
 * File: \go_concurrency\ch1\livelock\main.go                                  *
 * Project: Golang_Starting                                                    *
 * Created At: Wednesday, 2022/05/25 , 20:32:04                                *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Thursday, 2022/05/26 , 00:22:16                              *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package main

import (
	"bytes"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	cadence := sync.NewCond(&sync.Mutex{})

	go func() {
		for range time.Tick(1 * time.Millisecond) {
			cadence.Broadcast()
		}
	}()

	//takeStep 模拟所有对象之间的一个不变的节奏
	takeStep := func() {
		cadence.L.Lock()
		cadence.Wait()
		cadence.L.Unlock()
	}

	// tryDir 允许一个人尝试向一个方向移动，并返回是否成功，dir，试图朝这个方向移动的人数
	tryDir := func(dirname string, dir *int32, out *bytes.Buffer) bool {
		fmt.Fprintf(out, " %v", dirname)
		atomic.AddInt32(dir, 1)
		takeStep()

		if atomic.LoadInt32(dir) == 1 {
			fmt.Fprint(out, ". Success!")
			return true
		}

		takeStep()
		atomic.AddInt32(dir, -1) //这里的人意识到不能向这个方向走而放弃
		return false
	}

	var left, right int32

	tryLeft := func(out *bytes.Buffer) bool { return tryDir("left", &left, out) }
	tryRight := func(out *bytes.Buffer) bool { return tryDir("right", &right, out) }

	walk := func(walking *sync.WaitGroup, name string) {
		var out bytes.Buffer
		defer func() { fmt.Println(out.String()) }()

		defer walking.Done()

		fmt.Fprintf(&out, "%v is trying to scoot:", name)

		for i := 0; i < 5; i++ { //对尝试次数进行了限制，以便此程序能结束，在一个有活锁的程序中，可能没有这个限制
			if tryLeft(&out) || tryRight(&out) { //首先，这个人会试图向左走，如果失败了，他们会尝试向右走。
				return
			}
		}
		fmt.Fprintf(&out, "\n%v tosses hers hands up in exasperation!", name)
	}

	var peopleInHallway sync.WaitGroup //这个变量为程序提供了一个等待直到两个人都能够相互通过或放弃的方式。
	peopleInHallway.Add(2)
	go walk(&peopleInHallway, "Alice")
	go walk(&peopleInHallway, "Babara")
	peopleInHallway.Wait()
}
