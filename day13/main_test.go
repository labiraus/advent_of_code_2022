package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
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
				"[1,1,3,1,1]",
				"[1,1,5,1,1]",
				"",
				"[[1],[2,3,4]]",
				"[[1],4]",
				"",
				"[9]",
				"[[8,7,6]]",
				"",
				"[[4,4],4,4]",
				"[[4,4],4,4,4]",
				"",
				"[7,7,7,7]",
				"[7,7,7]",
				"",
				"[]",
				"[3]",
				"",
				"[[[]]]",
				"[[]]",
				"",
				"[1,[2,[3,[4,[5,6,7]]]],8,9]",
				"[1,[2,[3,[4,[5,6,0]]]],8,9]",
			},
			expected: 13,
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

			n := Dataset{}
			n.build(stream)

			output := n.eval()
			if output != test.expected {
				t.Errorf("expected: %v got: %v", test.expected, output)
			}
		})
	}
}

func Test_build(t *testing.T) {
	tests := []struct {
		name         string
		register     Dataset
		input        string
		expectedItem Item
	}{
		{name: "basic",
			register: Dataset{},
			input:    "[1,1,3,1,1]",
			expectedItem: Item{Data: []Item{
				{Value: 1, Data: []Item{}},
				{Value: 1, Data: []Item{}},
				{Value: 3, Data: []Item{}},
				{Value: 1, Data: []Item{}},
				{Value: 1, Data: []Item{}},
			}},
		},
		{name: "basic1",
			register:     Dataset{},
			input:        "[]",
			expectedItem: Item{Data: []Item{}, Empty: true},
		},
		{name: "basic2",
			register:     Dataset{},
			input:        "[[[]]]",
			expectedItem: Item{Data: []Item{{Data: []Item{{Data: []Item{}, Empty: true}}}}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actualItem := build(test.input)
			if !cmp.Equal(actualItem, test.expectedItem) {
				t.Errorf(cmp.Diff(test.expectedItem, actualItem))
				t.Errorf("\nexpected: %+v\ngot:      %+v", test.expectedItem, actualItem)
			}
		})
	}
}
