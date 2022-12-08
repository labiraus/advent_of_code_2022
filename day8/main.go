package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type grid struct {
	layout [][]tree
}
type tree struct {
	h    int
	nMax int
	sMax int
	wMax int
	eMax int
}

func main() {
	f, err := os.Open("day8/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	lines := make(chan string)
	n := grid{}
	go func() {
		for scanner.Scan() {
			text := scanner.Text()
			lines <- text
		}
		close(lines)
	}()
	n.build(lines)
	fmt.Println(n.eval())
}

func (n *grid) build(lines <-chan string) {
	var err error
	row := 0
	for line := range lines {
		n.layout = append(n.layout, make([]tree, len(line)))
		for col, char := range line {
			n.layout[row][col].h, err = strconv.Atoi(string(char))
			if err != nil {
				log.Fatal(err)
			}
			if col > 0 {
				n.layout[row][col].wMax = max(n.layout[row][col-1].wMax, n.layout[row][col-1].h)
			}
			if row > 0 {
				n.layout[row][col].nMax = max(n.layout[row-1][col].nMax, n.layout[row-1][col].h)
			}
		}
		row++
	}
}

func (n *grid) eval() (int, int) {
	sum := 0
	maxView := 0
	for row := len(n.layout) - 1; row >= 0; row-- {
		for col := len(n.layout[0]) - 1; col >= 0; col-- {

			if row < len(n.layout)-1 {
				n.layout[row][col].sMax = max(n.layout[row+1][col].sMax, n.layout[row+1][col].h)
			}

			if col < len(n.layout[0])-1 {
				n.layout[row][col].eMax = max(n.layout[row][col+1].eMax, n.layout[row][col+1].h)
			}
			// visibile if on the edge or visible
			if row == 0 || row == len(n.layout)-1 ||
				col == 0 || col == len(n.layout[0])-1 ||
				n.layout[row][col].visible() {
				sum++
			}
			view := n.calculateView(row, col)
			if view > maxView {
				maxView = view
			}
		}
	}
	return sum, maxView
}

func (n *grid) calculateView(row int, col int) int {
	wDist := 0
	if n.layout[row][col].h > n.layout[row][col].wMax {
		wDist = col
	} else {
		for i := col - 1; i >= 0; i-- {
			wDist++
			if n.layout[row][col].h <= n.layout[row][i].h {
				break
			}
		}
	}
	if wDist == 0 {
		return 0
	}

	eDist := 0
	if n.layout[row][col].h > n.layout[row][col].eMax {
		eDist = len(n.layout[0]) - col - 1
	} else {
		for i := col + 1; i < len(n.layout[0]); i++ {
			eDist++
			if n.layout[row][col].h <= n.layout[row][i].h {
				break
			}
		}
	}
	if eDist == 0 {
		return 0
	}

	nDist := 0
	if n.layout[row][col].h > n.layout[row][col].nMax {
		nDist = row
	} else {
		for i := row - 1; i >= 0; i-- {
			nDist++
			if n.layout[row][col].h <= n.layout[i][col].h {
				break
			}
		}
	}
	if nDist == 0 {
		return 0
	}

	sDist := 0
	if n.layout[row][col].h > n.layout[row][col].sMax {
		sDist = len(n.layout) - row - 1
	} else {
		for i := row + 1; i < len(n.layout); i++ {
			sDist++
			if n.layout[row][col].h <= n.layout[i][col].h {
				break
			}
		}
	}
	if sDist == 0 {
		return 0
	}

	return wDist * eDist * nDist * sDist
}

func (t *tree) visible() bool {
	return t.h > t.eMax || t.h > t.wMax || t.h > t.nMax || t.h > t.sMax
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
