/*
 * File: \src\ch9\memo1\main.go                                                *
 * Project: Golang_Starting                                                    *
 * Created At: Friday, 2022/05/20 , 12:52:02                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Friday, 2022/05/20 , 22:29:35                                *
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

type Func func(key string) (any, error)

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

func New(f Func) *Memo {
	memo := &Memo{make(chan request)}
	go memo.server(f)

	return memo
}

//并发安全
func (memo *Memo) Get(key string) (any, error) {
	response := make(chan result)
	memo.requests <- request{key, response}
	res := <-response

	return res.value, res.err
}

func (memo *Memo) Close() {
	close(memo.requests)
}

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)

	for req := range memo.requests {
		e := cache[req.key]

		if e == nil {
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(f, req.key)
		}

		go e.deliver(req.response)
	}
}

func (e *entry) call(f Func, key string) {
	e.res.value, e.res.err = f(key)

	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	<-e.ready

	response <- e.res
}
