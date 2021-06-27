package ex03

import (
	"fmt"
	"os"
	"strings"
)

func EchoUsingLoop() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func EchoUsingStringJoin() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
