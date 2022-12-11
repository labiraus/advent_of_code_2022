package main

import (
	"testing"
)

func Test_eval(t *testing.T) {
	tests := []struct {
		name     string
		register register
		input    []string
		expected int
	}{
		{name: "golden",
			register: register{cycle: 1, X: 1},
			input: []string{
				"addx 15",
				"addx -11",
				"addx 6",
				"addx -3",
				"addx 5",
				"addx -1",
				"addx -8",
				"addx 13",
				"addx 4",
				"noop",
				"addx -1",
				"addx 5",
				"addx -1",
				"addx 5",
				"addx -1",
				"addx 5",
				"addx -1",
				"addx 5",
				"addx -1",
				"addx -35",
				"addx 1",
				"addx 24",
				"addx -19",
				"addx 1",
				"addx 16",
				"addx -11",
				"noop",
				"noop",
				"addx 21",
				"addx -15",
				"noop",
				"noop",
				"addx -3",
				"addx 9",
				"addx 1",
				"addx -3",
				"addx 8",
				"addx 1",
				"addx 5",
				"noop",
				"noop",
				"noop",
				"noop",
				"noop",
				"addx -36",
				"noop",
				"addx 1",
				"addx 7",
				"noop",
				"noop",
				"noop",
				"addx 2",
				"addx 6",
				"noop",
				"noop",
				"noop",
				"noop",
				"noop",
				"addx 1",
				"noop",
				"noop",
				"addx 7",
				"addx 1",
				"noop",
				"addx -13",
				"addx 13",
				"addx 7",
				"noop",
				"addx 1",
				"addx -33",
				"noop",
				"noop",
				"noop",
				"addx 2",
				"noop",
				"noop",
				"noop",
				"addx 8",
				"noop",
				"addx -1",
				"addx 2",
				"addx 1",
				"noop",
				"addx 17",
				"addx -9",
				"addx 1",
				"addx 1",
				"addx -3",
				"addx 11",
				"noop",
				"noop",
				"addx 1",
				"noop",
				"addx 1",
				"noop",
				"noop",
				"addx -13",
				"addx -19",
				"addx 1",
				"addx 3",
				"addx 26",
				"addx -30",
				"addx 12",
				"addx -1",
				"addx 3",
				"addx 1",
				"noop",
				"noop",
				"noop",
				"addx -9",
				"addx 18",
				"addx 1",
				"addx 2",
				"noop",
				"noop",
				"addx 9",
				"noop",
				"noop",
				"noop",
				"addx -1",
				"addx 2",
				"addx -37",
				"addx 1",
				"addx 3",
				"noop",
				"addx 15",
				"addx -21",
				"addx 22",
				"addx -6",
				"addx 1",
				"noop",
				"addx 2",
				"addx 1",
				"noop",
				"addx -10",
				"noop",
				"noop",
				"addx 20",
				"addx 1",
				"addx 2",
				"addx 2",
				"addx -6",
				"addx -11",
				"noop",
				"noop",
				"noop",
			},
			expected: 13140,
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

			n := register{X: 1, cycle: 1}
			n.eval(stream)
			output := n.strength
			if output != test.expected {
				t.Errorf("expected: %v got: %v", test.expected, output)
			}
		})
	}
}

func Test_add(t *testing.T) {
	tests := []struct {
		name     string
		register register
		input    int
		expected int
	}{
		{
			name:     "test1",
			register: register{cycle: 1, X: 1, strength: 1},
			input:    3,
			expected: 1,
		},
		{
			name:     "test2",
			register: register{cycle: 19, X: 1, strength: 1},
			input:    3,
			expected: 21,
		},
		{
			name:     "test3",
			register: register{cycle: 18, X: 1, strength: 1},
			input:    3,
			expected: 1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.register.addx(test.input)
			if test.register.strength != test.expected {
				t.Errorf("expected: %v got: %v", test.expected, test.register.strength)
			}
		})
	}
}
