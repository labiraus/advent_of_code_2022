package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/beefsack/go-astar"
)

type step int32

const (
	down step = iota
	level
	up
)

type dataset struct {
	layout *layout
	startX int
	startY int
	endX   int
	endY   int
}

func main() {
	f, err := os.Open("day12/input.txt")

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

	r := dataset{}
	r.build(lines)
	minDist := 100000.0
	for x, line := range r.layout.w {
		for y, t := range line {
			if t.height == 0 {
				dist := r.eval(x, y)
				if dist > 0 && minDist > dist {
					minDist = dist
				}
			}
		}
	}
	fmt.Println(minDist)
}

func (t *dataset) build(lines <-chan string) {
	t.layout = &layout{make([][]*tile, 0)}
	x := 0
	for line := range lines {
		t.layout.w = append(t.layout.w, make([]*tile, len(line)))
		for y, val := range line {
			if val == 'S' {
				t.startX = x
				t.startY = y
				val = 'a'
			}
			if val == 'E' {
				t.endX = x
				t.endY = y
				val = 'z'
			}
			height := int(val - 'a')
			t.layout.w[x][y] = &tile{height: height, X: x, Y: y, layout: t.layout}
		}
		x++
	}
}

func (t *dataset) eval(startX, startY int) float64 {
	t1 := t.layout.tile(startX, startY)
	t2 := t.layout.tile(t.endX, t.endY)
	path, distance, found := astar.Path(t1, t2)
	if !found {
		return 0
	}
	for _, to := range path {
		to.(*tile).path = true
	}
	return distance
}

func (d *dataset) print() {
	for _, line := range d.layout.w {
		for _, t := range line {
			if d.startX == t.X && d.startY == t.Y {
				fmt.Print("S")
			} else if d.endX == t.X && d.endY == t.Y {
				fmt.Print("E")
			} else if t.path {
				fmt.Print("*")
			} else {
				fmt.Print(string(rune(t.height + int('a'))))
			}
		}
		fmt.Println()
	}
}

func (d *dataset) printTouched() {
	for _, line := range d.layout.w {
		for _, t := range line {
			if d.startX == t.X && d.startY == t.Y {
				fmt.Print("S")
			} else if d.endX == t.X && d.endY == t.Y {
				fmt.Print("E")
			} else if t.touched {
				fmt.Print("*")
			} else {
				fmt.Print(string(rune(t.height + int('a'))))
			}
		}
		fmt.Println()
	}
}
