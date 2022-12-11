package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type register struct {
	cycle    int
	X        int
	strength int
}

func main() {
	f, err := os.Open("day10/input.txt")

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

	r := register{X: 1, cycle: 1}
	r.eval(lines)
	fmt.Println()
	fmt.Println(r.strength)
}

func (r *register) eval(lines <-chan string) {
	for line := range lines {
		instructions := strings.Split(line, " ")
		if instructions[0] == "noop" {
			r.noop()
		} else {
			val, err := strconv.Atoi(instructions[1])
			if err != nil {
				panic(err)
			}
			r.addx(val)
		}
	}
}

func (r *register) noop() {
	r.duringCycle()
	r.cycle++
}

func (r *register) addx(val int) {
	// during first cycle
	r.duringCycle()
	// end of first cycle
	r.cycle++
	// during second cycle
	r.duringCycle()
	// end of second cycle
	r.X += val
	r.cycle++
}

func (r *register) duringCycle() {
	pos := r.cycle % 40
	if pos == 1 {
		fmt.Println()
	}
	if pos <= r.X+2 && pos >= r.X {
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}
	if (r.cycle-20)%40 == 0 {
		r.strength += r.cycle * r.X
	}
}
