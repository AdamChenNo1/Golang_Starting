package main

import (
	"fmt"
	"reflect"
)

type Strings interface {
	~string
}

func main() {
	p("1")
	// fmt.Println(less("sd", "nil"))
}

func p[T Strings](a T) {
	fmt.Printf("%#v\n", reflect.ValueOf(a))
	fmt.Printf("%T\n", reflect.ValueOf(a))
	fmt.Printf("%T\n", reflect.TypeOf(a))
}
func less[T Strings](a, b T) bool {
	switch reflect.TypeOf(a).Kind() {
	case 1:
		return false
	default:
		return true
	}
}
