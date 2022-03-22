/**
Write a non-recursive version of comma, using bytes.Buffer instead of string concat
*/

package main

import (
	"bytes"
	"fmt"
	"os"
)

func comma(s string) string {
	n := len(s)

	if n <= 3 {
		return s
	}

	var buf bytes.Buffer

	var commaIndexes []int

	var x int
	for i := n - 1; i > 0; i-- {
		if x == 2 {
			commaIndexes = append(commaIndexes, i)
			x = 0
		}
		x++
	}

	for i, r := range s {
		if contains(commaIndexes, i) {
			buf.WriteByte(',')
		}
		buf.WriteRune(r)
	}

	return buf.String()
}

func contains(values []int, i int) bool {
	for _, value := range values {
		if value == i {
			return true
		}
	}
	return false
}

func main() {
	for _, arg := range os.Args[1:] {
		fmt.Println(comma(arg))
	}
}
