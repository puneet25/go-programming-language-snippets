package main

import "strings"

func isAnagram(s1 string, s2 string) bool {
	s1 = strings.ReplaceAll(s1, " ", "")
	s2 = strings.ReplaceAll(s2, " ", "")

	s1 = strings.ToLower(s1)
	s2 = strings.ToLower(s2)

	if len(s1) != len(s2) {
		return false
	}

	m := make(map[rune]int)

	for _, key := range s1 {
		m[key]++
	}

	for _, key := range s2 {
		value := m[key]

		if value == 0 {
			return false
		}

		m[key]--
	}

	for _, value := range m {
		if value > 0 {
			return false
		}
	}

	return true
}
