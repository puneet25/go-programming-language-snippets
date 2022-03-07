/* Modify fetch all to print its output to a file so it can be examined*/
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string) // creates a channel of strings using make

	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine using `go`, calls fetch asynchronously
	}

	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from ch
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // writing to channel ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body) // reads the body and discards it by writing to the ioutil.Discard output stream
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()

	ch <- fmt.Sprint("%.2f %7d %s", secs, nbytes, url)
}
