package main

/*
 Modify fetch to add the prefix http:// to each argument url if it is missing.
*/
import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		response, err := http.Get(completeURL(url))
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetching: %v", err)
			os.Exit(1)
		}

		_, err = io.Copy(os.Stdout, response.Body)
		response.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "copying: %v", err)
			os.Exit(1)
		}
	}
}

func completeURL(url string) string {
	if (strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://")) == false {
		url = "http://" + url
	}

	return url
}
