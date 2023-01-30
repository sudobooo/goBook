// Echo_v2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	result, separator := "", ""
	for _, arg := range os.Args[1:] {
		result += separator + arg
		separator = " "
	}

	fmt.Println(result)
}
