package main

import (
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
				"Sensor at x=2, y=18: closest beacon is at x=-2, y=15",
				"Sensor at x=9, y=16: closest beacon is at x=10, y=16",
				"Sensor at x=13, y=2: closest beacon is at x=15, y=3",
				"Sensor at x=12, y=14: closest beacon is at x=10, y=16",
				"Sensor at x=10, y=20: closest beacon is at x=10, y=16",
				"Sensor at x=14, y=17: closest beacon is at x=10, y=16",
				"Sensor at x=8, y=7: closest beacon is at x=2, y=10",
				"Sensor at x=2, y=0: closest beacon is at x=2, y=10",
				"Sensor at x=0, y=11: closest beacon is at x=2, y=10",
				"Sensor at x=20, y=14: closest beacon is at x=25, y=17",
				"Sensor at x=17, y=20: closest beacon is at x=21, y=22",
				"Sensor at x=16, y=7: closest beacon is at x=15, y=3",
				"Sensor at x=14, y=3: closest beacon is at x=15, y=3",
				"Sensor at x=20, y=1: closest beacon is at x=15, y=3",
			},
			expected: 56000011,
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

			d := buildDataset()
			d.build(stream)
			x, y := d.plot(20, 20)
			if x*4000000+y != test.expected {
				t.Errorf("expected: %v got: %v", test.expected, x*y)
			}
		})
	}
}
