package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type direction int

const (
	R direction = iota
	L
	U
	D
)

var moves = map[string]direction{
	"R": R,
	"L": L,
	"U": U,
	"D": D,
}

type rope struct {
	posX           []int
	posY           []int
	tailPositions  map[int]map[int]bool
	tailPositions2 map[int]map[int]bool
}

func main() {
	f, err := os.Open("day9/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	lines := make(chan string)
	go func() {
		for scanner.Scan() {
			text := scanner.Text()
			lines <- text
		}
		close(lines)
	}()
	r := rope{
		tailPositions: make(map[int]map[int]bool),
		posX:          make([]int, 10),
		posY:          make([]int, 10),
	}
	fmt.Println(r.eval(lines))
}

func (r *rope) eval(lines <-chan string) int {
	for line := range lines {
		commands := strings.Split(line, " ")
		dir := commands[0]
		count, err := strconv.Atoi(commands[1])
		if err != nil {
			panic(err)
		}
		for i := 0; i < count; i++ {
			r.move(moves[dir])
		}
	}

	sum := 0
	for _, rows := range r.tailPositions {
		sum += len(rows)
	}
	return sum
}

func (r *rope) move(dir direction) {
	switch dir {
	case U:
		r.posX[0]++
	case D:
		r.posX[0]--
	case R:
		r.posY[0]++
	case L:
		r.posY[0]--
	}

	for i := 1; i < len(r.posX); i++ {
		r.posX[i], r.posY[i] = moveTail(r.posX[i-1], r.posY[i-1], r.posX[i], r.posY[i])
	}

	//r.updateTailPos(1)
	r.updateTailPos(len(r.posX) - 1)
}

func (r *rope) updateTailPos(tail int) {
	row, ok := r.tailPositions[r.posX[tail]]
	if !ok {
		row = make(map[int]bool)
		r.tailPositions[r.posX[tail]] = row
	}
	row[r.posY[tail]] = true
}

func moveTail(frontx, fronty, backx, backy int) (int, int) {
	distx := frontx - backx
	disty := fronty - backy
	switch distx {
	case 2:
		backx++
	case -2:
		backx--

	case 1:
		if Abs(disty) == 2 {
			backx++
		}
	case -1:
		if Abs(disty) == 2 {
			backx--
		}
	}

	switch disty {
	case 2:
		backy++
	case -2:
		backy--

	case 1:
		if Abs(distx) == 2 {
			backy++
		}
	case -1:
		if Abs(distx) == 2 {
			backy--
		}
	}
	return backx, backy
}

func (r *rope) print() {
	maxX := 0
	maxY := 0
	for val, positions := range r.tailPositions {
		if val > maxX {
			maxX = val
		}
		for pos := range positions {
			if pos > maxY {
				maxY = pos
			}
		}
	}
	for x := maxX; x >= 0; x-- {
		for y := 0; y <= maxY; y++ {
			if _, ok := r.tailPositions[x]; ok && r.tailPositions[x][y] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
