/*
 * File: \go_concurrency\ch3\pipline\pipline.go                                *
 * Project: Golang_Starting                                                    *
 * Created At: Friday, 2022/05/27 , 18:22:53                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Friday, 2022/05/27 , 18:58:24                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package pipline

// 这个函数会重复传给它的值，直到告诉它停止
func Repeat(done <-chan any, values ...any) <-chan any {
	valueStream := make(chan any)
	go func() {
		defer close(valueStream)

		for {
			for _, v := range values {
				select {
				case <-done:
					return
				case valueStream <- v:
				}
			}
		}
	}()

	return valueStream
}

// 一个重复调用函数fn的生成器
func RepeatFn(done <-chan any, fn func() any) <-chan any {
	valueStream := make(chan any)

	go func() {
		defer close(valueStream)

		for {
			select {
			case <-done:
				return
			case valueStream <- fn():
			}
		}
	}()

	return valueStream
}

// 这个pipeline stage只会从其传入的valueStream中取出第一个num项目，然后退出
func Take(done <-chan any, valueStream <-chan any, num int) <-chan any {
	takeStream := make(chan any)
	go func() {
		defer close(takeStream)
		for i := 0; i < num; i++ {
			select {
			case <-done:
				return
			case takeStream <- <-valueStream:
			}
		}
	}()
	return takeStream
}

func ToString(done <-chan interface{}, valueStream <-chan interface{}) <-chan string {
	stringStream := make(chan string)
	go func() {
		defer close(stringStream)
		for v := range valueStream {
			// fmt.Printf("%T", v)
			select {
			case <-done:
				return
			case stringStream <- v.(string):
			}
		}
	}()

	return stringStream
}

func OrDone(done, c <-chan any) <-chan any {
	valStream := make(chan any)
	go func() {
		defer close(valStream)
		for {
			select {
			case <-done:
				return
			case v, ok := <-c:
				if !ok {
					return
				}
				select {
				case valStream <- v:
				case <-done:
				}
			}
		}
	}()

	return valStream
}

func Tee(done <-chan any, in <-chan any) (_, _ <-chan any) {
	out1, out2 := make(chan any), make(chan any)

	go func() {
		defer close(out1)
		defer close(out2)

		for val := range OrDone(done, in) {
			var out1, out2 = out1, out2 // 将要使用out1和out2的本地版本，所以隐藏这些变量
			for i := 0; i < 2; i++ {    // 为确保两者都写入，将执行select语句的两次迭代：每个出站一个channel
				select { // 使用一条select语句，以便不阻塞的写入out1和out2。
				case <-done:
				case out1 <- val: // 一旦写入了channel，将其影副本设置为nil，以便进一步阻塞写入，而另一个channel可以继续。
					out1 = nil
				case out2 <- val:
					out2 = nil
				}
			}
		}
	}()

	return out1, out2
}

func Bridge(done <-chan any, chanStream <-chan <-chan any) <-chan any {
	valStream := make(chan any) // 将返回bridge中的所有值的channel

	go func() {
		defer close(valStream)

		for { // 这个循环负责从chanStream中提取channel并将其提供给嵌套循环来使用
			var stream <-chan any
			select {
			case maybeStream, ok := <-chanStream:
				if !ok {
					return
				}
				stream = maybeStream
			case <-done:
				return
			}

			for val := range OrDone(done, stream) {
				// 该循环负责读取已经给出的channel中的值，并将这些值重复到valStream中
				// 当当前正在循环的流关闭时，从执行从此channel读取的循环中跳出，并继续循环的下一次迭代，
				// 选择要读取的channel，这为我们提供了一个不间断的结果值的流。
				select {
				case valStream <- val:
				case <-done:
				}
			}
		}
	}()

	return valStream
}
