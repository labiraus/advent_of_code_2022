package main

import (
	"fmt"
	"testing"
)

func Test_eval(t *testing.T) {
	tests := []struct {
		name     string
		register monkeys
		input    []string
		expected int64
	}{
		{name: "golden",
			register: monkeys{},
			input: []string{
				"				Monkey 0:",
				"				Starting items: 79, 98",
				"				Operation: new = old * 19",
				"				Test: divisible by 23",
				"				  If true: throw to monkey 2",
				"				  If false: throw to monkey 3",
				"			  ",
				"			  Monkey 1:",
				"				Starting items: 54, 65, 75, 74",
				"				Operation: new = old + 6",
				"				Test: divisible by 19",
				"				  If true: throw to monkey 2",
				"				  If false: throw to monkey 0",
				"			  ",
				"			  Monkey 2:",
				"				Starting items: 79, 60, 97",
				"				Operation: new = old * old",
				"				Test: divisible by 13",
				"				  If true: throw to monkey 1",
				"				  If false: throw to monkey 3",
				"			  ",
				"			  Monkey 3:",
				"				Starting items: 74",
				"				Operation: new = old + 3",
				"				Test: divisible by 17",
				"				  If true: throw to monkey 0",
				"				  If false: throw to monkey 1",
			},
			expected: 2_713_310_158,
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

			n := monkeys{}
			n.build(stream)

			for i := 1; i <= 10_000; i++ {
				n.round()
				if i%1_000 == 0 || i == 1 || i == 20 {
					fmt.Println("== After round ", i, "==")

					for num, m := range n.troop {
						fmt.Printf("Monkey %v inspected items %v times.\n", num, m.inspections)

					}
					fmt.Println()
				}
			}
			output := n.total()
			if output != test.expected {
				t.Errorf("expected: %v got: %v", test.expected, output)
			}
		})
	}
}
