/*
	fetchall fetches many URL's contents concurrently. It discards the responses but reports the size and elapsed time for each one
*/
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
	/* 	make is used to initialize an object of type slice, map and channel only.
	   	channel is a communication mechanism that allows one goroutine to pass values of a specified type to another goroutine
		a goroutine is a concurrent function execution
	*/
	ch := make(chan string) // creates a channel of strings using make

	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine using `go`, calls fetch asynchronously
	}

	for range os.Args[1:] {
		/*
			when one goroutine sends or receives, it blocks until another goroutine attempts a corresponding receive or send.
			here, in this example each fetch sends a value (ch<- expression) on the channel ch and main receives all of them (<-ch).
			meaning main is blocked until each fetch does not send the value
		*/
		fmt.Println(<-ch) // receive from ch
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

/*
	chan<- ????
	This is known as channel direction

	The rule of thumb: Arrow shows if the data is going into (output) or going out of (input) channel. No arrow is general purpose channel.
	chan <-          writing to channel (output channel)
	<- chan          reading from channel (input channel)
	chan             read from or write to channel (input/output channel)
*/
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

	ch <- fmt.Sprint("%.2fs %7d %s", secs, nbytes, url)
}
