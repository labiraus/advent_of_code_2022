package main

import (
	"testing"
)

func Test_eval(t *testing.T) {
	tests := []struct {
		name     string
		register dataset
		input    []string
		expected float64
	}{
		{name: "golden",
			register: dataset{},
			input: []string{
				"Sabqponm",
				"abcryxxl",
				"accszExk",
				"acctuvwj",
				"abdefghi",
			},
			expected: 31,
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

			data := dataset{}
			data.build(stream)
			output := data.eval(data.startX, data.startY)
			if output != test.expected {
				t.Errorf("expected: %v got: %v", test.expected, output)
			}
			data.print()
		})
	}
}
