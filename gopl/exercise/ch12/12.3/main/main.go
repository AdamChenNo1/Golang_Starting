/*
 * File: \src\ch12\sexpr\main.go                                               *
 * Project: Golang_Starting                                                    *
 * Created At: Monday, 2022/05/23 , 20:52:16                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Wednesday, 2022/05/25 , 20:07:02                             *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package main

import (
	"fmt"
	sexpr "go_start/gopl/exercise/ch12/12.3"
)

func main() {
	type Movie struct {
		Title, Subtitle string
		Year            int
		Color           bool
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
		score           float32
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
		score: 9.4,
	}
	b, err := sexpr.MarshalIndent(strangelove)

	if err == nil {
		fmt.Println(string(b))
	}

	type st struct {
		a string
		b int
	}
	m := map[st][]st{
		{"s", 5}: {
			{},
			{"54", 78},
			{"4", 780},
		},
		{"j", 15}: {
			{},
			{"754", 78},
			{"45", 780},
		},
	}

	b, err = sexpr.MarshalIndent(m)

	if err == nil {
		fmt.Println(string(b))
	}
}
