/*
	A minimal server that returns the path component of the URL used to access the server.
	e.g. http://localhost:8000/hello -> "/hello"

	Most of the work is done by Go library functions
*/

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// connects a handler function to incoming URLs whose path begins with "/", meaning this will listen to all URLs
	http.HandleFunc("/", handler)
	// starts a server listening for incoming requests on port 8000
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// echoes the path component of the requested URL
func handler(w http.ResponseWriter, r *http.Request) {
	// Using fmt.Fprintf to write to an http.ResponseWriter representing the web browser
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path) // use %q: to double-quote strings
}

/*
	How to start a server in the background on a Mac. Add an ampersand go
	go run server/server1.go &

	We can now access the server from a web browser by typing localhost:8000
*/
