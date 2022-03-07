//  Here is how we can use net/http’s get call to fetch the content of each specified url
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Prints the content found at each specified url
func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		// struct resp has a Body field. contains response as a readable stream
		// ioutil is used to read the stream
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
