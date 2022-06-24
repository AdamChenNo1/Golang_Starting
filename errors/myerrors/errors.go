/*
 * File: /myerrors/errors.go                                                   *
 * Project: errors                                                             *
 * Created At: Monday, 2022/06/20 , 12:36:16                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Monday, 2022/06/20 , 18:43:17                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package main

import (
	"fmt"

	"github.com/elchn/errors"
	code "github.com/marmotedu/sample-code"
)

func main() {
	e := errors.New("hello")
	e1 := errors.WrapC(e, code.ErrBind, "hi")
	fmt.Printf("%T\n", e)
	fmt.Printf("%v\n", e)
	fmt.Printf("%-v\n", e)
	fmt.Printf("%#-v\n", e)
	fmt.Printf("%+v\n", e)
	fmt.Printf("%#-v\n", e)
	fmt.Printf("%#+v\n", e)
	fmt.Println(e.Error())
	fmt.Println(e)
	fmt.Println("----------------------------------")

	fmt.Printf("%T\n", e1)
	fmt.Printf("%v\n", e1)
	fmt.Printf("%-v\n", e1)
	fmt.Printf("%+v\n", e1)
	fmt.Printf("%#-v\n", e1)
	fmt.Printf("%#+v\n", e1)
	fmt.Println(e1.Error())
	fmt.Println(e1.Error())
	fmt.Println(errors.ParseCoder(e1).String())

	fmt.Println(e1)
	// bytes, _ := json.MarshalIndent(myError, "", " ")
	// fmt.Println(string(bytes))
}
