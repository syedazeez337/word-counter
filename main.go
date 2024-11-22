package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func count(r io.Reader) int {
	// A scanner is used to read text from a Reader
	scanner := bufio.NewScanner(r)
	// Define the scanner split type to words
	scanner.Split(bufio.ScanWords)
	wc := 0

	for scanner.Scan() {
		wc++
	}

	return wc
}

func main() {
	counter := count(os.Stdin)
	fmt.Println(counter)
}
