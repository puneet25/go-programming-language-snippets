package main

import "testing"

func TestAnagram(t *testing.T) {

	var tests = []struct {
		s1   string
		s2   string
		want bool
	}{
		{"a", "a", true},
		{"a", "b", false},
		{"abc", "cba", true},
		{"aabb", "abab", true},
		{"aAbBcC", "abcABC", true},
		{"aAbBcC", "abcABCD", false},
	}

	for _, test := range tests {
		if got := isAnagram(test.s1, test.s2); got != test.want {
			t.Errorf("anagram(%s, %s) = %t, want %t", test.s1, test.s2, got, test.want)
		}
	}
}
