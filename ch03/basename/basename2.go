/**
Removes any prefix of s that looks like a system path
e.g. "a/b/c.go" -> "c"
	 "c.d.go" -> "c.d"
*/
package main

import (
	"fmt"
	"strings"
)

// This function uses standard library function
func basename(s string) string {
	// Discard last `\` and everything before
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

func main() {
	strings := []string{
		"a/b/c.go",
		"c.d.go",
		"abc",
	}
	for _, arg := range strings {
		fmt.Println(basename(arg))
	}
}
