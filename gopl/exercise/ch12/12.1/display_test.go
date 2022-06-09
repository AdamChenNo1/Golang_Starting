/*
 * File: \exercise\ch12\12.1\display_test.go                                   *
 * Project: Golang_Starting                                                    *
 * Created At: Monday, 2022/05/23 , 19:58:13                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Monday, 2022/05/23 , 20:02:21                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package display

func ExampleDisplay() {
	type s struct {
		f1 string
		f2 int
	}
	m := map[s]string{
		{f1: "Jack", f2: 100}: "passed",
		{f1: "Tom", f2: 50}:   "failed",
	}
	Display("m", m)
	//Output:
	// jk
}
