/*
 * File: \src\ch9\memo1\main.go                                                *
 * Project: Golang_Starting                                                    *
 * Created At: Friday, 2022/05/20 , 12:52:02                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Saturday, 2022/05/21 , 02:24:28                              *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package memo

type entry struct {
	res   result
	ready chan struct{} //res赋值后关闭
}

type Func struct {
	done chan struct{}
	f    func(key string) (any, error)
}

func NewFunc(f func(key string) (any, error)) *Func {
	return &Func{
		done: make(chan struct{}),
		f:    f,
	}
}

// type Func func(key string, done chan struct{}) (any, error)

type result struct {
	value any
	err   error
}

//request是一条请求消息，key需要用Func来调用
type request struct {
	key      string
	response chan<- result //客户端需要单个result
}

//the result of calling Func was stored in Memo
type Memo struct {
	requests chan request
}

func New(f *Func) *Memo {
	memo := &Memo{make(chan request)}
	go memo.server(f)

	return memo
}

//并发安全
//key url
//memoize 是否缓存结果
func (memo *Memo) Get(f *Func, key string, memoize bool) (any, error) {
	select {
	case <-f.done:
	default:
		if !memoize {
			close(f.done)
		}
	}

	response := make(chan result)

	memo.requests <- request{key, response}
	res := <-response

	return res.value, res.err

}

func (memo *Memo) Close() {
	close(memo.requests)
}

// func cancelled() bool {
// 	select {
// 	case <-done:
// 		return true
// 	default:
// 		return false
// 	}
// }

func (f Func) cancelled() bool {
	select {
	case <-f.done:
		return true
	default:
		return false
	}
}

func (memo *Memo) server(f *Func) {
	cache := make(map[string]*entry)

	for req := range memo.requests {
		if f.cancelled() {
			e := &entry{ready: make(chan struct{})}
			e.call(f, req.key)
			e.deliver(req.response)
		} else {
			e := cache[req.key]

			if e == nil {
				e = &entry{ready: make(chan struct{})}
				cache[req.key] = e
				go e.call(f, req.key)
			}

			go e.deliver(req.response)
		}
	}

}

func (e *entry) call(f *Func, key string) {
	e.res.value, e.res.err = f.f(key)

	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	<-e.ready

	response <- e.res
}
