package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	rock     = 1
	paper    = 2
	scissors = 3
	loss     = 0
	draw     = 3
	win      = 6
)

var (
	playValue = map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	playType = map[string]int{
		"A": rock,
		"B": paper,
		"C": scissors,
		"X": rock,
		"Y": paper,
		"Z": scissors,
	}

	winType = map[string]int{
		"X": loss,
		"Y": draw,
		"Z": win,
	}
)

func main() {

	f, err := os.Open("day2/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)

	score1 := 0
	score2 := 0
	for scanner.Scan() {
		text := scanner.Text()

		result, err := eval(text)
		if err != nil {
			panic(err)
		}
		score1 += result

		result2, err := eval2(text)
		if err != nil {
			panic(err)
		}
		score2 += result2
	}
	fmt.Println(score1)
	fmt.Println(score2)
}

func eval(input string) (int, error) {
	score := 0
	s := strings.Split(input, " ")
	myPlay, ok := playType[s[1]]
	if !ok {
		return 0, fmt.Errorf("could not parse play [%v]", s[1])
	}
	theirPlay, ok := playType[s[0]]
	if !ok {
		return 0, fmt.Errorf("could not parse play [%v]", s[0])
	}
	switch {
	case theirPlay == myPlay:
		// draw
		score += draw
	case theirPlay == myPlay+2 || theirPlay == myPlay-1:
		score += win
	}
	score += playValue[s[1]]

	return score, nil
}

func eval2(input string) (int, error) {
	s := strings.Split(input, " ")
	winning, ok := winType[s[1]]
	if !ok {
		return 0, fmt.Errorf("could not parse win [%v]", s[1])
	}
	theirPlay, ok := playType[s[0]]
	if !ok {
		return 0, fmt.Errorf("could not parse play [%v]", s[0])
	}
	myPlay := theirPlay
	switch winning {
	case win:
		myPlay = theirPlay + 1
	case loss:
		myPlay = theirPlay - 1
	}
	if myPlay > scissors {
		myPlay = rock
	} else if myPlay < rock {
		myPlay = scissors
	}

	return myPlay + winning, nil
}
