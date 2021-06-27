package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadingFromStandardInput() {
	// create a map named counts, which has string as key and int as value
	// make() can only be used to initialize slices, maps, and channels, and that, unlike the new() function, make() does not return a pointer.
	counts := make(map[string]int) // another way is  map[string]int{}

	countLines(os.Stdin, counts)

	// for loop for map behaving similar to Swift, i.e. lhs is key, rhs is value
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func ReadingFromNamedFiles() {
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				// %v - prints any value
				fmt.Printf("error opening named files: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	// for loop for map behaving similar to Swift, i.e. lhs is key, rhs is value
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// A map is reference to the data structure that is created by make. It means here the function receives the copy of the reference(passed by reference)
func countLines(f *os.File, counts map[string]int) {
	// bufio package's Scanner is reading input; easiest way to process input that comes in lines
	// to stop input in terminal, if reading from standard input; press CTRL+D; in real world, it will step at end of file
	input := bufio.NewScanner(f)

	// each call to input.Scan() reads the next line; returns either true or false; result can be retrived from input.Text()
	for input.Scan() {
		// shortcut statement to increase the value stored against a particular key;
		// note in Go, how we do not need to initialize the value first to 0 before incrementing it;
		// this is because in Go everyone has a default value, so fist time we write counts[input.Text()], it evaluates to 0
		counts[input.Text()]++
	}
}

func main() {
	ReadingFromNamedFiles()
}
