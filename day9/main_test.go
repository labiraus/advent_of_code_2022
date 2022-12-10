package main

import (
	"testing"
)

func Test_eval(t *testing.T) {
	tests := []struct {
		name     string
		rope     rope
		input    []string
		expected int
	}{{
		name: "golden",
		rope: rope{},
		input: []string{
			"R 4",
			"U 4",
			"L 3",
			"D 1",
			"R 4",
			"D 1",
			"L 5",
			"R 2"},
		expected: 13,
	}}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			stream := make(chan string)
			go func() {
				for _, line := range test.input {
					stream <- line
				}
				close(stream)
			}()
			r := rope{tailPositions: make(map[int]map[int]bool)}
			output := r.eval(stream)
			r.print()
			if output != test.expected {
				t.Errorf("expected: %v got: %v\n%+v", test.expected, output, r.tailPositions)
			}
		})
	}
}
