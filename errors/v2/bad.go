/*
 * File: \v2\bad.go                                                            *
 * Project: errors                                                             *
 * Created At: Thursday, 2022/06/16 , 13:51:07                                 *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Monday, 2022/06/20 , 18:21:08                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package main

import (
	"log"

	"github.com/elchn/errors"
)

func main() {
	if err := funcA(); err != nil {
		log.Fatal("failed calling func: ", err)
		return
	}

	log.Println("succeeded calling func")
}

func funcA() error {
	if err := funcB(); err != nil {
		return errors.Wrap(err, "failed calling funcB")
	}

	return errors.New("calling func error")
}

func funcB() error {
	return errors.New("calling func error")
}
