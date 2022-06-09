/*
 * File: \src\ch5\outline\v3\main\outline_test.go                              *
 * Project: gopl                                                               *
 * Created At: Wednesday, 2021/08/11 , 20:40:38                                *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Wednesday, 2022/06/8 , 23:42:13                              *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package main

import (
	"fmt"
	ch3 "go_start/gopl/src/ch5/outline/v3"
	"os"

	"golang.org/x/net/html"
)

func ExampleForEachNode() {
	filename := "C:\\Users\\Elon\\Desktop\\golang\\Golang_Starting\\src\\outline\\v3\\index.html"

	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "open %s： %v", filename, err)
		return
	}

	doc, err := html.Parse(file)

	if err != nil {
		fmt.Fprintf(os.Stderr, "parse HTML %s： %v", filename, err)
		return
	}

	ch3.ForEachNode(doc, ch3.StartElement, ch3.EndElement)
	// Output:
	// <html>
	//   <head/>
	//   <body>
	//     <button>
	//       <a/ href="sfdsfg.com">
	//     </button>
	//     <div>
	//       Hello world!
	//     </div>
	//     <img/ src="golang.jpeg">
	//   </body>
	// </html>

}
