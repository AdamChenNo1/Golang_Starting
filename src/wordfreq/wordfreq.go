package ch5

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func WordFreq(in io.Reader) map[string]uint {
	freq := make(map[string]uint)

	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		freq[scanner.Text()]++
		// fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "read input: %v", err)
	}

	return freq
}
