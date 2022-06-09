package main

import (
	"fmt"
	ch2 "go_start/src/outline/v3"
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

	ch2.ForEachNode(doc, ch2.StartElement, ch2.EndElement)
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
