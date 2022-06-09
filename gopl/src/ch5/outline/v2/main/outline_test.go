/*
 * File: \src\ch5\outline\v2\main\outline_test.go                              *
 * Project: gopl                                                               *
 * Created At: Wednesday, 2021/08/11 , 19:51:38                                *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Wednesday, 2022/06/8 , 23:41:06                              *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package main

import (
	"fmt"
	ch2 "go_start/gopl/src/ch5/outline/v2"
	"os"

	"golang.org/x/net/html"
)

func ExampleForEachNode() {
	filename := "C:\\Users\\Elon\\Desktop\\golang\\Golang_Starting\\src\\outline\\v2\\index.html"

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

	ch2.ForEachNode(doc, ch2.StartElement, ch2.EndElement)
	// Output:
	// <html>
	//   <head>
	//   </head>
	//   <body>
	//     <button>
	//       <a>
	//       </a>
	//     </button>
	//     <img>
	//     </img>
	//   </body>
	// </html>

}
