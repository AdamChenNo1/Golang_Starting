/*
 * File: \exercise\ch9\9.4\main.go                                             *
 * Project: Golang_Starting                                                    *
 * Created At: Saturday, 2022/05/21 , 14:23:32                                 *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Saturday, 2022/05/21 , 18:30:56                              *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package main

func main() {
	g1 := make(chan struct{}, 1)
	g1 <- struct{}{}
	for {
		go func(in chan struct{}) {
			v := <-in
			in <- v
		}(g1)
	}
}
