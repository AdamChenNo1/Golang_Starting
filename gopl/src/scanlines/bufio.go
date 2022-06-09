/*
 * File: \src\scanlines\bufio.go                                               *
 * Project: Golang_Starting                                                    *
 * Created At: Tuesday, 2022/04/19 , 17:07:30                                  *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Tuesday, 2022/04/19 , 18:30:31                               *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package main

import (
	"fmt"
)

func main() {
	// sc := bufio.NewScanner(os.Stdin)
	var x int = 32
	fmt.Printf("%c", x)

	n, _ := fmt.Scanln(&x)
	if n > 0 {
		fmt.Println(x)
	}
}
