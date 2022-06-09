/*
 * File: \src\ch9\memo1\main.go                                                *
 * Project: Golang_Starting                                                    *
 * Created At: Friday, 2022/05/20 , 12:52:02                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Friday, 2022/05/20 , 14:55:29                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package memo

import "sync"

type entry struct {
	res   result
	ready chan struct{} //res赋值后关闭
}

//the result of calling Func was stored in Memo
type Memo struct {
	f     Func
	mu    sync.Mutex
	cache map[string]*entry
}

type Func func(key string) (any, error)

type result struct {
	value any
	err   error
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: map[string]*entry{}}
}

//并发安全
func (memo *Memo) Get(key string) (any, error) {
	memo.mu.Lock()
	e := memo.cache[key]

	if e == nil {
		//The first time when key was accessed,computing the value and broadcasting it
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()
		e.res.value, e.res.err = memo.f(key)

		close(e.ready) //when value is set,close the channel ready so each goroutine accessing it will not be blocked now
	} else {
		memo.mu.Unlock()
		<-e.ready
	}

	return e.res.value, e.res.err
}
