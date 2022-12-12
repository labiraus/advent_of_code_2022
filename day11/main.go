package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type opp int

const (
	mult opp = iota
	add
	square
)

type monkey struct {
	items        []int
	operation    opp
	operationVal int
	testVal      int
	trueMonkey   int
	falseMonkey  int
	inspections  int64
}

type monkeys struct {
	primeSum int
	troop    []monkey
}

func main() {
	f, err := os.Open("day11/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	lines := make(chan string)
	go func() {
		for scanner.Scan() {
			lines <- scanner.Text()
		}
		close(lines)
	}()

	r := monkeys{}
	r.build(lines)
	for i := 0; i < 10_000; i++ {
		r.round()
	}
	fmt.Println(r.total())
}

func (t *monkeys) build(lines <-chan string) {
	factors := []int{}
	m := monkey{items: make([]int, 0)}
	for line := range lines {
		line = strings.TrimSpace(line)
		instructions := strings.Split(line, ":")
		if len(line) == 0 {
			t.troop = append(t.troop, m)
			m = monkey{items: make([]int, 0)}
		} else {
			vals := strings.Split(strings.TrimSpace(instructions[1]), " ")
			switch strings.TrimSpace(instructions[0]) {
			case "Starting items":
				for _, valString := range vals {
					val, err := strconv.Atoi(strings.TrimRight(valString, ","))
					if err != nil {
						panic(err)
					}
					m.items = append(m.items, val)
				}
			case "Operation":
				if vals[4] == "old" {
					m.operation = square
					continue
				}
				val, err := strconv.Atoi(vals[4])
				if err != nil {
					panic(err)
				}
				m.operationVal = val
				if vals[3] == "*" {
					m.operation = mult
				} else {
					m.operation = add
				}

			case "Test":
				val, err := strconv.Atoi(vals[2])
				if err != nil {
					panic(err)
				}
				m.testVal = val
				factors = append(factors, val)

			case "If true":
				val, err := strconv.Atoi(vals[3])
				if err != nil {
					panic(err)
				}
				m.trueMonkey = val

			case "If false":
				val, err := strconv.Atoi(vals[3])
				if err != nil {
					panic(err)
				}
				m.falseMonkey = val
			}
		}
	}
	t.troop = append(t.troop, m)

	primeSum := 1
	for _, factor := range factors {
		primeSum *= factor
	}
	t.primeSum = primeSum
}

func (t *monkeys) round() {
	for i := 0; i < len(t.troop); i++ {
		m := t.troop[i]
		for _, item := range m.items {
			t.troop[i].inspections++
			// Inspect
			switch m.operation {
			case mult:
				item *= m.operationVal
			case square:
				item *= item
			case add:
				item += m.operationVal
			}
			item %= t.primeSum

			// Relax
			//item /= 3

			// Test
			throw := item%m.testVal == 0
			// Throw
			if throw {
				t.troop[m.trueMonkey].items = append(t.troop[m.trueMonkey].items, item)
			} else {
				t.troop[m.falseMonkey].items = append(t.troop[m.falseMonkey].items, item)
			}
		}
		t.troop[i].items = make([]int, 0)
	}
}

func (t *monkeys) total() int64 {
	troop := t.troop
	sort.SliceStable(troop, func(i int, j int) bool {
		return troop[i].inspections > troop[j].inspections
	})
	return troop[0].inspections * troop[1].inspections
}
