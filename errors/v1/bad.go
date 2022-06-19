/*
 * File: \bad.go                                                               *
 * Project: errors                                                             *
 * Created At: Thursday, 2022/06/16 , 13:36:26                                 *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Thursday, 2022/06/16 , 13:39:47                              *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package main

import (
	"fmt"
	"log"
)

func main() {
	if err := funcA(); err != nil {
		log.Fatal("failed calling func:", err)
		return
	}

	log.Println("succeeded calling func")
}

func funcA() error {
	if err := funcB(); err != nil {
		return err
	}

	return fmt.Errorf("calling func error")
}

func funcB() error {
	return fmt.Errorf("calling func error")
}
