package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Dataset struct {
	Items  map[int]map[int]Material
	GrainX int
	GrainY int
	Bottom int
	Left   int
	Right  int
}

type Material int

const (
	Air Material = iota
	Rock
	Sand
)

func main() {
	f, err := os.Open("day14/input.txt")

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

	d := Dataset{Items: make(map[int]map[int]Material), Left: 500, Right: 500}
	d.build(lines)
	fmt.Println(d.eval())
}

func (d *Dataset) build(lines <-chan string) {
	for line := range lines {
		line = strings.TrimSpace(line)
		paths := strings.Split(line, " -> ")
		for i := 0; i < len(paths)-1; i++ {
			startx, starty, endx, endy := parsePath(paths, i)
			if starty == endy {
				for x := startx; x <= endx; x++ {
					d.setValue(x, starty, Rock)
				}
			} else {
				for y := starty; y <= endy; y++ {
					d.setValue(startx, y, Rock)
				}
			}
			if endy > d.Bottom {
				d.Bottom = endy + 1
			}
			if startx <= d.Left {
				d.Left = startx - 1
			}
			if endx >= d.Right {
				d.Right = endx + 1
			}
		}
	}
}

func parsePath(paths []string, i int) (int, int, int, int) {
	start := strings.Split(paths[i], ",")
	end := strings.Split(paths[i+1], ",")
	startx, err := strconv.Atoi(start[0])
	if err != nil {
		panic(err)
	}
	starty, err := strconv.Atoi(start[1])
	if err != nil {
		panic(err)
	}
	endx, err := strconv.Atoi(end[0])
	if err != nil {
		panic(err)
	}
	endy, err := strconv.Atoi(end[1])
	if err != nil {
		panic(err)
	}
	if starty > endy {
		starty, endy = endy, starty
	}
	if startx > endx {
		startx, endx = endx, startx
	}
	return startx, starty, endx, endy
}

func (d *Dataset) eval() int {
	sum := 0
	for d.dropGrain() {
		sum++
	}

	return sum
}

func (d *Dataset) dropGrain() bool {
	d.GrainX = 500
	d.GrainY = 0
	if d.isSolid(d.GrainX, d.GrainY) {
		return false
	}

	for d.moveGrain() {
		if d.GrainY == d.Bottom {
			d.setValue(d.GrainX, d.GrainY, Sand)
			break
		}
	}
	return true
}

func (d *Dataset) moveGrain() bool {
	switch {
	case !d.isSolid(d.GrainX, d.GrainY+1):
		// Drop down
		d.GrainY++
		return true
	case !d.isSolid(d.GrainX-1, d.GrainY+1):
		// Drop down left
		d.GrainX--
		d.GrainY++
		return true
	case !d.isSolid(d.GrainX+1, d.GrainY+1):
		// Drop down right
		d.GrainX++
		d.GrainY++
		return true
	default:
		// Settle
		d.setValue(d.GrainX, d.GrainY, Sand)
		return false
	}
}

func (d *Dataset) isSolid(x, y int) bool {
	_, ok := d.Items[x][y]
	return ok
}

func (d Dataset) String() string {
	out := ""
	for y := 0; y < d.Bottom; y++ {
		for x := d.Left; x <= d.Right; x++ {
			char := "."
			if mat, ok := d.Items[x][y]; ok {
				switch mat {
				case Rock:
					char = "#"
				case Sand:
					char = "O"
				default:
					char = "X"
				}
			}
			out += char
		}
		out += "\n"
	}
	return out
}

func (d *Dataset) setValue(x, y int, mat Material) {
	if _, ok := d.Items[x]; !ok {
		d.Items[x] = make(map[int]Material)
	}
	d.Items[x][y] = mat
}
