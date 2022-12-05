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

	f, err := os.Open("day5/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	stacks := setup(scanner)

	// for _, stack := range stacks {
	// 	fmt.Println(stack)
	// }

	for scanner.Scan() {
		text := scanner.Text()
		//fmt.Println(text)
		act(text, stacks)
	}

	for _, stack := range stacks {
		fmt.Print(string(stack[len(stack)-1]))
	}
}

func setup(scanner *bufio.Scanner) [][]string {
	initialState := []string{}
	// scan through initial stack state until empty line hit
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 {
			break
		}
		initialState = append(initialState, text)
	}

	// initialise stacks
	stacks := [][]string{}
	width := len(initialState[0])/4 + 1
	// fmt.Println("width: ", width)
	for i := 0; i < width; i++ {
		stacks = append(stacks, []string{})
	}

	// work through initial state from bottom to top
	for i := len(initialState) - 2; i >= 0; i-- {
		for j := 0; j < width; j++ {
			pos := (j * 4) + 1
			char := []rune(initialState[i])[pos]
			if char != ' ' {
				stacks[j] = append(stacks[j], string(char))
			}
		}
	}

	return stacks
}

func act(text string, stacks [][]string) {
	action := strings.Split(text, " ")

	moveCount, err := strconv.Atoi(action[1])
	if err != nil {
		panic(err)
	}
	origin, err := strconv.Atoi(action[3])
	if err != nil {
		panic(err)
	}
	destination, err := strconv.Atoi(action[5])
	if err != nil {
		panic(err)
	}
	// part 1 code
	// for i := 0; i < moveCount; i++ {
	// 	moveCrate(origin-1, destination-1, stacks)
	// }

	// fmt.Print(moveCount, " ", stacks[origin-1], stacks[destination-1], " -> ")
	moveStack(origin-1, destination-1, moveCount, stacks)
	// fmt.Print(stacks[origin-1], stacks[destination-1], "\n")

}

func moveCrate(origin int, destination int, stacks [][]string) {
	//fmt.Print(stacks[origin], " ", stacks[destination], "\n")
	crate := stacks[origin][len(stacks[origin])-1]
	stacks[origin] = stacks[origin][:len(stacks[origin])-1]
	stacks[destination] = append(stacks[destination], crate)
}

func moveStack(origin int, destination int, moveCount int, stacks [][]string) {

	crates := stacks[origin][len(stacks[origin])-moveCount : len(stacks[origin])]
	// fmt.Printf("(%v) ", crates)
	stacks[origin] = stacks[origin][:len(stacks[origin])-moveCount]
	stacks[destination] = append(stacks[destination], crates...)
}
