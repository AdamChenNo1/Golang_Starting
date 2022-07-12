/*
 * File: /dependency_injection/main.go                                         *
 * Project: go-demo                                                            *
 * Created At: Sunday, 2022/06/26 , 12:52:19                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Sunday, 2022/06/26 , 12:55:23                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package main

import "os"

func main() {
	Greet(os.Stdout, "Elodie\n")
}
