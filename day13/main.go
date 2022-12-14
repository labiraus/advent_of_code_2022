package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/google/go-cmp/cmp"
)

type Dataset struct {
	Items []Item
}

type Item struct {
	Data  []Item
	Value int
	Empty bool
}

type outcome int

const (
	ordered outcome = iota
	disorderd
	unknown
)

func main() {
	f, err := os.Open("day13/input.txt")

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

	d := Dataset{}
	d.build(lines)
	fmt.Println(d.eval())
	i1 := Item{Data: []Item{{Value: 2}}}
	i2 := Item{Data: []Item{{Value: 6}}}
	d.Items = append(d.Items, i1)
	d.Items = append(d.Items, i2)
	sort.Slice(d.Items, func(i, j int) bool {
		return compare(d.Items[i], d.Items[j]) != disorderd
	})
	sum := 1
	for i, item := range d.Items {
		if cmp.Equal(item, i1) || cmp.Equal(item, i2) {
			sum *= i + 1
		}
	}
	fmt.Println(sum)
}

func (d *Dataset) build(lines <-chan string) {
	for line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		item := build(line)
		d.Items = append(d.Items, item)
		// fmt.Println("e: " + line)
		// fmt.Println("g: " + item.print())
	}
}

func (t *Dataset) eval() int {
	sum := 0
	for i := 0; i < len(t.Items); i += 2 {
		if compare(t.Items[i], t.Items[i+1]) != disorderd {
			sum += i/2 + 1
		}
	}

	return sum
}

func compare(left, right Item) outcome {
	// Handle if one list is empty
	switch {
	case left.Empty && !right.Empty:
		return ordered
	case !left.Empty && right.Empty:
		return disorderd
	case left.Empty && right.Empty:
		return unknown
	}

	// Compare values
	if len(left.Data) == 0 && len(right.Data) == 0 {
		switch {
		case left.Value < right.Value:
			return ordered
		case left.Value == right.Value:
			return unknown
		case left.Value > right.Value:
			return disorderd
		}
	}

	// Convert left value to list
	if len(left.Data) == 0 {
		return compare(Item{Data: []Item{{Value: left.Value}}}, right)
	}

	// Convert right value to list
	if len(right.Data) == 0 {
		return compare(left, Item{Data: []Item{{Value: right.Value}}})
	}

	// Compare lists
	for i, rightItem := range right.Data {
		if i >= len(left.Data) {
			return ordered
		}
		out := compare(left.Data[i], rightItem)
		if out != unknown {
			return out
		}
	}

	// Right list ran out of items first
	if len(right.Data) < len(left.Data) {
		return disorderd
	}

	return unknown
}

func build(line string) Item {
	item, _, _ := buildData(line[1:])
	return item
}

func buildData(line string) (Item, string, bool) {
	if line[0] == ']' {
		return Item{Empty: true, Data: make([]Item, 0)}, line[1:], true
	}

	out := Item{Data: make([]Item, 0)}
	for len(line) > 0 {
		commaIndex := strings.Index(line, ",")
		openIndex := strings.Index(line, "[")
		closeIndex := strings.Index(line, "]")

		switch {
		case commaIndex == 0:
			line = line[commaIndex+1:]
		case commaIndex >= 0 && (commaIndex < openIndex || openIndex == -1) && (commaIndex < closeIndex || closeIndex == -1):
			val, err := strconv.Atoi(line[:commaIndex])
			if err != nil {
				panic(err)
			}
			out.Data = append(out.Data, Item{Data: make([]Item, 0), Value: val})
			line = line[commaIndex+1:]

		case openIndex >= 0 && (openIndex < closeIndex || closeIndex == -1):
			var newItem Item
			ok := false
			newItem, line, ok = buildData(line[openIndex+1:])
			if ok {
				out.Data = append(out.Data, newItem)
			}

		case closeIndex > 0:
			val, err := strconv.Atoi(line[:closeIndex])
			if err != nil {
				panic(err)
			}
			out.Data = append(out.Data, Item{Data: make([]Item, 0), Value: val})
			return out, line[closeIndex+1:], true

		default:
			return out, line[closeIndex+1:], true
		}
	}

	return out, line, true
}

func (i *Item) print() string {
	switch {
	case i.Empty:
		return "[]"
	case len(i.Data) == 0:
		return strconv.Itoa(i.Value)
	default:
		vals := []string{}
		for _, item := range i.Data {
			vals = append(vals, item.print())
		}
		return "[" + strings.Join(vals, ",") + "]"
	}
}
