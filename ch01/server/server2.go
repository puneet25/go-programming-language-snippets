/*
	Adding a counter
*/

package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex // to acquire lock and prevent race conditions
var count int 

func main() {
	// two handlers
	// 1. invokes all URLs
	// 2. invokes request to /count
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// behind the scene server runs these requests in a go routine to process requests asynchronously, hence the Lock and Unlock
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

/*
	How to start a server in the background on a Mac. Add an ampersand go
	go run server/server1.go &

	We can now access the server from a web browser by typing localhost:8000
*/
