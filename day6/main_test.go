package main

import (
	"testing"
)

func Test_gold(t *testing.T) {
	tests := []struct {
		name       string
		expected   int
		input      string
		markerSize int
	}{
		{
			name:       "example 1:1",
			input:      "bvwbjplbgvbhsrlpgdmjqwftvncz",
			expected:   5,
			markerSize: 4,
		},
		{
			name:       "example 1:2",
			input:      "nppdvjthqldpwncqszvftbrmjlhg",
			expected:   6,
			markerSize: 4,
		},
		{
			name:       "example 1:3",
			input:      "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			expected:   10,
			markerSize: 4,
		},
		{
			name:       "example 1:4",
			input:      "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			expected:   11,
			markerSize: 4,
		},

		{
			name:       "example 2:1",
			input:      "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			expected:   19,
			markerSize: 14,
		},
		{
			name:       "example 2:2",
			input:      "bvwbjplbgvbhsrlpgdmjqwftvncz",
			expected:   23,
			markerSize: 14,
		},
		{
			name:       "example 2:3",
			input:      "nppdvjthqldpwncqszvftbrmjlhg",
			expected:   23,
			markerSize: 14,
		},
		{
			name:       "example 2:4",
			input:      "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			expected:   29,
			markerSize: 14,
		},
		{
			name:       "example 2:5",
			input:      "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			expected:   26,
			markerSize: 14,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			d := data{
				marker:     []string{},
				markerSize: test.markerSize,
			}

			for _, char := range test.input {
				ok := d.eval(string(char))
				if ok {
					break
				}
			}
			if test.expected != d.i {
				t.Errorf("expected: %v, got: %v\n%+v", test.expected, d.i, d)
			}
		})
	}
}
