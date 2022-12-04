package main

import "testing"

func Test_eval(t *testing.T) {
	tests := []struct {
		name     string
		expected bool
		input    string
	}{{
		"2-88,13-89",
		false,
		"2-88,13-89",
	}}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			actual := eval(test.input)
			if actual != test.expected {
				t.Errorf("expected: %v, got: %v", test.expected, actual)
			}
		})
	}
}
