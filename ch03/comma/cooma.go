/**
Inserts comma every three places
"12345" -> "12,345"
*/

package main

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	return comma(s[:n-3]) + "," + comma(s[n-3:])
}
