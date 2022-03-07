package main

import (
	"fmt"
	"os"

	"example.com/greetings"
)

func main() {
	var s, sep string
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)

	fmt.Println(greetings.Hello(s))
}
