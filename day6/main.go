package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type data struct {
	i          int
	marker     []string
	markerSize int
}

func main() {
	f, err := os.Open("day6/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	reader := bufio.NewReader(f)

	d := data{
		marker:     make([]string, 0, 4),
		markerSize: 14,
	}

	for {
		if c, _, err := reader.ReadRune(); err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		} else {
			found := d.eval(string(c))
			if found {
				break
			}
		}
	}
	fmt.Println(d.i)
}

func (d *data) eval(char string) bool {
	d.i++
	for pos, oldChar := range d.marker {
		if oldChar == char {
			d.marker = append(d.marker[pos+1:], char)
			return false
		}

		if pos == d.markerSize-2 {
			return true
		}
	}
	d.marker = append(d.marker, char)

	return false
}
