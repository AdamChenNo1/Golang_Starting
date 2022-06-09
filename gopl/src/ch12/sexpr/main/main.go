/*
 * File: \src\ch12\sexpr\main.go                                               *
 * Project: Golang_Starting                                                    *
 * Created At: Monday, 2022/05/23 , 20:52:16                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Monday, 2022/05/23 , 21:06:35                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package main

import (
	"fmt"
	"go_start/src/ch12/sexpr"
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
	b, err := sexpr.Marshal(strangelove)

	if err == nil {
		fmt.Println(string(b))
	}
}
