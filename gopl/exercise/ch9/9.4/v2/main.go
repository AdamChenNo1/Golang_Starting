/*
 * File: \exercise\ch9\9.4\v2\main.go                                          *
 * Project: Golang_Starting                                                    *
 * Created At: Saturday, 2022/05/21 , 18:49:31                                 *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Monday, 2022/05/23 , 01:01:48                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package main

import (

	// "runtime/trace"

	"net/http"
	_ "net/http/pprof"
	"sync"
)

// func init() {
// 	runtime.SetMutexProfileFraction(1)
// }

func main() {
	// traceout, err := os.Create("./trace.out")
	// if err != nil {
	// 	panic(err)
	// }
	// defer traceout.Close()

	// err = trace.Start(traceout)
	// if err != nil {
	// 	panic(err)
	// }

	// defer trace.Stop()

	out := make(chan struct{})
	var wg sync.WaitGroup

	var f func(in chan struct{})

	wg.Add(1)
	f = func(in chan struct{}) {
		wg.Add(1)
		defer wg.Done()
		v := <-in
		out <- v
		go f(out)
	}
	go f(out)
	out <- struct{}{}

	http.ListenAndServe("0.0.0.0:6060", nil)

	wg.Wait()

}
