package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	// Idiomatic Test setup: https://gobyexample.com/testing
	var tests = []struct {
		name   string
		expect string
	}{
		{"TestHello 1", "Hello to Romén's Golang boilerplate project!"},
	}

	for _, tt := range tests {
		testname := tt.name
		t.Run(testname, func(t *testing.T) {
			ans := hello()
			if ans != tt.expect {
				t.Errorf("got %s, want %s", ans, tt.expect)
			}
		})
	}
}
