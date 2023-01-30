// Fetch prints the content found at each specified URL.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "https://") {
			url = "https://" + url
		}

		resp, err := http.Get(url) //nolint:gosec // because

		if err != nil {
			_, printErr := fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			if printErr != nil {
				return
			}

			os.Exit(1)
		}

		_, err = io.Copy(os.Stdout, resp.Body)
		fmt.Println(resp.Status)
		errClose := resp.Body.Close()

		if errClose != nil {
			return
		}

		if err != nil {
			_, printErr := fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			if printErr != nil {
				return
			}

			os.Exit(1)
		}
	}
}
