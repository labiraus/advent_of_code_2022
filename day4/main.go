package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	f, err := os.Open("day4/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	sum := 0
	sum2 := 0
	for scanner.Scan() {
		text := scanner.Text()
		if eval(text) {
			sum++
		}
		if eval2(text) {
			sum2++
		}
	}
	fmt.Println(sum)
	fmt.Println(sum2)
}

func eval(text string) bool {
	elves := strings.Split(text, ",")
	elf1 := strings.Split(elves[0], "-")
	elf2 := strings.Split(elves[1], "-")
	elf1Min, _ := strconv.Atoi(elf1[0])
	elf1Max, _ := strconv.Atoi(elf1[1])
	elf2Min, _ := strconv.Atoi(elf2[0])
	elf2Max, _ := strconv.Atoi(elf2[1])
	if (elf1Min <= elf2Min && elf1Max >= elf2Max) ||
		(elf1Min >= elf2Min && elf1Max <= elf2Max) {
		return true
	}
	return false
}

func eval2(text string) bool {
	elves := strings.Split(text, ",")
	elf1 := strings.Split(elves[0], "-")
	elf2 := strings.Split(elves[1], "-")
	elf1Min, _ := strconv.Atoi(elf1[0])
	elf1Max, _ := strconv.Atoi(elf1[1])
	elf2Min, _ := strconv.Atoi(elf2[0])
	elf2Max, _ := strconv.Atoi(elf2[1])
	if elf1Min <= elf2Max && elf1Max >= elf2Min {
		return true
	}
	return false
}
