package main

import (
	"fmt"
	"testing"
)

func Test_eval(t *testing.T) {
	tests := []struct {
		name     string
		register Dataset
		input    []string
		expected int
	}{
		{name: "golden",
			register: Dataset{},
			input: []string{
				"498,4 -> 498,6 -> 496,6",
				"503,4 -> 502,4 -> 502,9 -> 494,9",
			},
			expected: 93,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			stream := make(chan string)
			go func() {
				for _, line := range test.input {
					stream <- line
				}
				close(stream)
			}()

			d := Dataset{Items: make(map[int]map[int]Material), Left: 500, Right: 500}
			d.build(stream)
			fmt.Println(d.String())

			output := d.eval()
			fmt.Println(d.String())
			if output != test.expected {
				t.Errorf("expected: %v got: %v", test.expected, output)
			}
		})
	}
}
