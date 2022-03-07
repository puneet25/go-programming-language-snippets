package main

/*
 The function call io.Copy(dst, src) reads from src and writes to dst.
 Use it instead if ioutilReadAll to copy the response body to os.Stdout without requiring a buffer large enough to hold the entire stream.
*/
import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		response, err := http.Get(url)
		if err != nil {
			fmt.Printf("Error: %v fetching %s", err, url)
			os.Exit(1)
		}

		_, err = io.Copy(os.Stdout, response.Body)
		response.Body.Close()

		if err != nil {
			fmt.Printf("Error: %v copying %s response to writer", err, url)
			os.Exit(1)
		}
	}
}
