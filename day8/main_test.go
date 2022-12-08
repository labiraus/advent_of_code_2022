package main

import (
	"testing"
)

func Test_eval(t *testing.T) {
	tests := []struct {
		name     string
		grid     grid
		input    []string
		expected int
	}{
		{
			"golden",
			grid{layout: [][]tree{
				{tree{h: 3, nMax: 0, sMax: 6, wMax: 0, eMax: 7}, tree{h: 0, nMax: 0, sMax: 5, wMax: 3, eMax: 7}, tree{h: 3, nMax: 0, sMax: 5, wMax: 3, eMax: 7}, tree{h: 7, nMax: 0, sMax: 9, wMax: 3, eMax: 3}, tree{h: 3, nMax: 0, sMax: 9, wMax: 7, eMax: 0}},
				{tree{h: 2, nMax: 3, sMax: 6, wMax: 0, eMax: 5}, tree{h: 5, nMax: 0, sMax: 5, wMax: 2, eMax: 5}, tree{h: 5, nMax: 3, sMax: 5, wMax: 5, eMax: 2}, tree{h: 1, nMax: 7, sMax: 9, wMax: 5, eMax: 2}, tree{h: 2, nMax: 3, sMax: 9, wMax: 5, eMax: 0}},
				{tree{h: 6, nMax: 3, sMax: 3, wMax: 0, eMax: 5}, tree{h: 5, nMax: 5, sMax: 5, wMax: 6, eMax: 3}, tree{h: 3, nMax: 5, sMax: 5, wMax: 6, eMax: 3}, tree{h: 3, nMax: 7, sMax: 9, wMax: 6, eMax: 2}, tree{h: 2, nMax: 3, sMax: 9, wMax: 6, eMax: 0}},
				{tree{h: 3, nMax: 6, sMax: 3, wMax: 0, eMax: 9}, tree{h: 3, nMax: 5, sMax: 5, wMax: 3, eMax: 9}, tree{h: 5, nMax: 5, sMax: 3, wMax: 3, eMax: 9}, tree{h: 4, nMax: 7, sMax: 9, wMax: 5, eMax: 9}, tree{h: 9, nMax: 3, sMax: 0, wMax: 5, eMax: 0}},
				{tree{h: 3, nMax: 6, sMax: 0, wMax: 0, eMax: 9}, tree{h: 5, nMax: 5, sMax: 0, wMax: 3, eMax: 9}, tree{h: 3, nMax: 5, sMax: 0, wMax: 5, eMax: 9}, tree{h: 9, nMax: 7, sMax: 0, wMax: 5, eMax: 0}, tree{h: 0, nMax: 9, sMax: 0, wMax: 9, eMax: 0}},
			}},
			[]string{
				"30373",
				"25512",
				"65332",
				"33549",
				"35390",
			},
			21,
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

			n := grid{layout: [][]tree{}}
			n.build(stream)
			output, _ := n.eval()
			if output != test.expected {
				t.Errorf("expected: %v got: %v", test.expected, output)
			}
			for i := 0; i < len(test.grid.layout); i++ {
				for j := 0; j < len(test.grid.layout[0]); j++ {
					if test.grid.layout[i][j] != n.layout[i][j] {
						t.Errorf("expected: %+v got: %+v at %v,%v", test.grid.layout[i][j], n.layout[i][j], i, j)
					}
				}
			}
		})
	}
}

func Test_calculateView(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected int
		row      int
		col      int
	}{
		{
			"golden 1",
			[]string{
				"30373",
				"25512",
				"65332",
				"33549",
				"35390",
			},
			4,
			1,
			2,
		},
		{
			"golden 2",
			[]string{
				"30373",
				"25512",
				"65332",
				"33549",
				"35390",
			},
			8,
			3,
			2,
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

			n := grid{layout: [][]tree{}}
			n.build(stream)
			n.eval()
			output := n.calculateView(test.row, test.col)
			if output != test.expected {
				t.Errorf("expected: %v got: %v", test.expected, output)
			}

		})
	}
}
