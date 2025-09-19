package main

import (
	"testing"
)

func TestFindRootDomain(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"stinascloud.ir.", "stinascloud.ir"},
		{"stinascloud.ir", "stinascloud.ir"},
		{"ae-01.stinascloud.ir.", "stinascloud.ir"},
		{"ae-01.stinascloud.ir", "stinascloud.ir"},
		{"sub.ae-01.stinascloud.ir.", "stinascloud.ir"},
		{"example.com.", "example.com"},
		{"www.example.com.", "example.com"},
		{"api.v1.example.com.", "example.com"},
	}

	for _, test := range tests {
		result := findRootDomain(test.input)
		if result != test.expected {
			t.Errorf("findRootDomain(%s) = %s; expected %s", test.input, result, test.expected)
		}
	}
}
