/*
 * File: \test\http_get\main.go                                                *
 * Project: Golang_Starting                                                    *
 * Created At: Thursday, 2022/05/19 , 16:12:26                                 *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Thursday, 2022/05/19 , 16:16:17                              *
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
	"net/http"
)

func main() {
	// url := "http://gopl.io"
	url := "http://baidu.com"

	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatal(resp.Status)
	}

	fmt.Println(resp.Body)
}
