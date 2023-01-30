// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				_, printErr := fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				if printErr != nil {
					return
				}
				continue
			}
			countLines(f, counts)
			err = f.Close()
			if err != nil {
				return
			}
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	// NOTE: ignoring potential errors from input.Err()
	for input.Scan() {
		counts[input.Text()]++
	}
}
