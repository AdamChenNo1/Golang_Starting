package bytecounter

import "fmt"

func ExampleByteCounter() {
	var c ByteCounter

	c.Write([]byte("hello"))

	fmt.Println(c)

	c = 0
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)
	// Output:
	//
	// 5
	// 12
}
