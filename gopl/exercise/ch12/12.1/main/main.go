/*
 * File: \src\ch12\format\main\main.go                                         *
 * Project: Golang_Starting                                                    *
 * Created At: Monday, 2022/05/23 , 16:45:21                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Wednesday, 2022/05/25 , 20:06:38                             *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package main

import (
	"fmt"
	display "go_start/gopl/exercise/ch12/12.1"
	"reflect"
)

func main() {

	type Movie struct {
		Title, Subtitle string
		Year            int
		Color           bool
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
	}

	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worring and Love the Bomb",
		Year:     1964,
		Color:    false,
		Actor: map[string]string{
			"Dr. Strangelove":            "PeterSellers",
			"Grp. Capt. Lionel Mandrake": "PeterSellers",
			"Pres. Merkin Muffley":       "PeterSellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen Jack D. Ripper":   "Sterling Hayden",
			`Maj. T.J "King" Kong`:       "Slim Pickens",
		},
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted ScreenPlay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}
	fmt.Println(display.Format(reflect.ValueOf(strangelove)))
	fmt.Println()

	display.Display("strangelove", strangelove)
	// display.Display("os.Stderr", os.Stderr)
	fmt.Println()
	fmt.Println(display.Format(reflect.ValueOf([3]int{1, 2, 3})))
	fmt.Println()
	arr1 := [3]struct {
		a string
		b int
	}{
		{a: "Jack", b: 1},
		{a: "Tom", b: 2},
		{a: "Mike", b: 3},
	}
	fmt.Println(display.Format(reflect.ValueOf(arr1)))
	fmt.Println()
	type grades struct {
		a string
		b int
	}
	arr2 := [3]grades{
		{a: "Jack", b: 1},
		{a: "Tom", b: 2},
		{a: "Mike", b: 3},
	}
	fmt.Println(display.Format(reflect.ValueOf(arr2)))
	fmt.Println()
	slice1 := []struct {
		a string
		b int
	}{
		{a: "Jack", b: 1},
		{a: "Tom", b: 2},
		{a: "Mike", b: 3},
	}
	fmt.Println(display.Format(reflect.ValueOf(slice1)))
	fmt.Println()
	slice2 := []grades{
		{a: "Jack", b: 1},
		{a: "Tom", b: 2},
		{a: "Mike", b: 3},
	}
	fmt.Println(display.Format(reflect.ValueOf(slice2)))

	fmt.Println()
	m := map[grades]string{
		{a: "Jack", b: 100}: "passed",
		{a: "Tom", b: 50}:   "failed",
	}
	fmt.Println(display.Format(reflect.ValueOf(m)))
}
