/**
Removes any prefix of s that looks like a system path
e.g. "a/b/c.go" -> "c"
	 "c.d.go" -> "c.d"
*/
package main

import (
	"fmt"
)

// This function does not use any standard library function
func basename(s string) string {
	// Discard last `\` and everything before
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
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
