package main

import (
	"fmt"
	ch5 "go_start/src/wordfreq"
	"os"
)

func main() {
	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "open %s: %v", filename, err)
	}

	wordfreq := ch5.WordFreq(file)

	for word, freq := range wordfreq {
		fmt.Printf("%s : %d\n", word, freq)
	}
}
