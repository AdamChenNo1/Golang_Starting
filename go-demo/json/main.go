/*
 * File: \test\json\main.go                                                    *
 * Project: Golang_Starting                                                    *
 * Created At: Tuesday, 2022/05/24 , 01:17:27                                  *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Tuesday, 2022/05/24 , 01:22:52                               *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {

	x := map[int]complex64{
		1: 1 + 2i,
		2: 2 + 3i,
	}
	y, _ := json.Marshal(x)
	s := strings.Builder{}
	s.Write(y)
	fmt.Println(s.String())
}
