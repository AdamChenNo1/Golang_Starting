package tempflag

import (
	"flag"
	"fmt"
)

func ExampleCelsiusFlag() {
	var temp = CelsiusFlag("temp", 20.0, "the temperature")

	flag.Parse()
	fmt.Println(*temp)
	// Output:
	// 20°C
}
