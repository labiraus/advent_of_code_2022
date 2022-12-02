package main

import "testing"

func Test_eval(t *testing.T) {
	tests := []struct {
		name     string
		expected int
		input    string
	}{
		{
			name:     "rock = rock",
			expected: 4,
			input:    "A X",
		},
		{
			name:     "rock < paper",
			expected: 1,
			input:    "B X",
		},
		{
			name:     "rock > scissors",
			expected: 7,
			input:    "C X",
		},

		{
			name:     "paper > rock",
			expected: 8,
			input:    "A Y",
		},
		{
			name:     "paper = paper",
			expected: 5,
			input:    "B Y",
		},
		{
			name:     "paper < scissors",
			expected: 2,
			input:    "C Y",
		},

		{
			name:     "scissors < rock",
			expected: 3,
			input:    "A Z",
		},
		{
			name:     "scissors > paper",
			expected: 9,
			input:    "B Z",
		},
		{
			name:     "scissors = scissors",
			expected: 6,
			input:    "C Z",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			actual, err := eval(test.input)
			if err != nil {
				t.Fatalf(err.Error())
			}
			if actual != test.expected {
				t.Errorf("expected: %v, got: %v", test.expected, actual)
			}
		})
	}
}

func Test_eval2(t *testing.T) {
	tests := []struct {
		name     string
		expected int
		input    string
	}{
		{
			name:     "rock loss = scissors",
			expected: 3,
			input:    "A X",
		},
		{
			name:     "paper loss = rock",
			expected: 1,
			input:    "B X",
		},
		{
			name:     "scissor loss = paper",
			expected: 2,
			input:    "C X",
		},

		{
			name:     "rock draw = rock",
			expected: 4,
			input:    "A Y",
		},
		{
			name:     "paper draw = paper",
			expected: 5,
			input:    "B Y",
		},
		{
			name:     "scissor draw = scissors",
			expected: 6,
			input:    "C Y",
		},

		{
			name:     "rock win = paper",
			expected: 8,
			input:    "A Z",
		},
		{
			name:     "paper win = scissors",
			expected: 9,
			input:    "B Z",
		},
		{
			name:     "scissor win = rock",
			expected: 7,
			input:    "C Z",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			actual, err := eval2(test.input)
			if err != nil {
				t.Fatalf(err.Error())
			}
			if actual != test.expected {
				t.Errorf("expected: %v, got: %v", test.expected, actual)
			}
		})
	}

}
