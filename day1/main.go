package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	f, err := os.Open("day1/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	elves := []int{}
	elf := 0
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			elves = append(elves, elf)
			elf = 0
		} else {
			cal, err := strconv.Atoi(text)
			if err != nil {
				fmt.Println(err)
			} else {
				elf += cal
			}
		}
	}

	sort.Ints(elves)
	fmt.Println(elves[len(elves)-1])
	fmt.Println(elves[len(elves)-1] + elves[len(elves)-2] + elves[len(elves)-3])
}
