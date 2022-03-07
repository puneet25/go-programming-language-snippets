package main

import (
	"testing"
)

func TestCompleteURL(t *testing.T) {
	var tests = []struct {
		url         string
		completeURL string
	}{
		{"gopl.io", "http://gopl.io"},
		{"http://gopl.io", "http://gopl.io"},
		{"", "http://"},
		{"http://", "http://"},
		{"https://gopl.io", "https://gopl.io"},
	}

	for _, test := range tests {
		if got := completeURL(test.url); got != test.completeURL {
			t.Errorf("completeURL(%s) = %s, completeURL %s", test.url, got, test.completeURL)
		}
	}
}
