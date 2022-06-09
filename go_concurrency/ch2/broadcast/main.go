/*
 * File: \go_concurrency\ch2\broadcast\main.go                                 *
 * Project: Golang_Starting                                                    *
 * Created At: Thursday, 2022/05/26 , 16:09:42                                 *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Thursday, 2022/05/26 , 16:22:23                              *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package main

import (
	"fmt"
	"sync"
)

func main() {

	type Button struct { // 定义了一个Button类型，它包含一个条件，Clicked
		Clicked *sync.Cond
	}

	button := Button{Clicked: sync.NewCond(&sync.Mutex{})}

	// 定义了一个便利构造函数，它允许注册函数处理来自条件的信号
	// 每个处理程序都在自己的goroutine上运行，并且订阅不会退出，直到goroutine被确认运行为止
	subscribe := func(c *sync.Cond, fn func()) { //
		var goroutineRunning sync.WaitGroup
		goroutineRunning.Add(1)

		go func() {
			goroutineRunning.Done()
			c.L.Lock()
			defer c.L.Unlock()

			c.Wait()
			fn()
		}()
		goroutineRunning.Wait()
	}

	// 创建一个WaitGroup，这只是为了确保程序在写入stdout之前不会退出
	var clickRegistered sync.WaitGroup

	clickRegistered.Add(3)

	// 注册一个处理程序，当单击按键时，它将模拟最大化按钮的窗口
	subscribe(button.Clicked, func() {
		fmt.Println("Maximizing window.")
		clickRegistered.Done()
	})

	// 注册一个处理程序，该处理程序在单击鼠标时模拟显示对话框
	subscribe(button.Clicked, func() {
		fmt.Println("Displaying annoying dialog box!")
		clickRegistered.Done()
	})

	// 接下来，模拟一个用户通过单击应用程序的按钮来单击鼠标按键
	subscribe(button.Clicked, func() {
		fmt.Println("Mouse clicked.")
		clickRegistered.Done()
	})

	// 为鼠标按键事件设置了一个处理程序，它反过来调用Cond上的Broadcast，让所有的处理程序都知道鼠标按键已经被单击了（更健壮的
	// 实现将首先检查它是否已经被抑制）
	button.Clicked.Broadcast()
	clickRegistered.Wait()
}
