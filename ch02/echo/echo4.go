/*
Echo4 prints its command line arguments.
It takes two optional flags;
-n: to omit the trailing new line
-s: to separate the output arguments by the contents of the string sep instead of thea default single space.
*/
package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing new line")
var sep = flag.String("sep", " ", "separator")

func main() {
	// Parse parses the command-line flags from os.Args[1:]. Must be called after all flags are defined and before flags are accessed by the program.
	flag.Parse()

	// Non flag arguments are available from flag.Args() as a slice of strings.
	fmt.Print(strings.Join(flag.Args(), *sep))
	if *n == false {
		fmt.Println()
	}
}

type Celsius float64
type Fahrenheit float64

func test() {

	var cValue Celsius = 25
	var fValue Fahrenheit = 25
	fmt.Println(cValue == fvalue)
}
