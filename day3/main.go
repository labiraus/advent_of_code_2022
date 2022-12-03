package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var priorities = map[rune]int{
	'a': 1,
	'b': 2,
	'c': 3,
	'd': 4,
	'e': 5,
	'f': 6,
	'g': 7,
	'h': 8,
	'i': 9,
	'j': 10,
	'k': 11,
	'l': 12,
	'm': 13,
	'n': 14,
	'o': 15,
	'p': 16,
	'q': 17,
	'r': 18,
	's': 19,
	't': 20,
	'u': 21,
	'v': 22,
	'w': 23,
	'x': 24,
	'y': 25,
	'z': 26,
	'A': 27,
	'B': 28,
	'C': 29,
	'D': 30,
	'E': 31,
	'F': 32,
	'G': 33,
	'H': 34,
	'I': 35,
	'J': 36,
	'K': 37,
	'L': 38,
	'M': 39,
	'N': 40,
	'O': 41,
	'P': 42,
	'Q': 43,
	'R': 44,
	'S': 45,
	'T': 46,
	'U': 47,
	'V': 48,
	'W': 49,
	'X': 50,
	'Y': 51,
	'Z': 52,
}

func main() {

	f, err := os.Open("day3/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	itemMap := make([]map[rune]bool, 3)
	sum1 := 0
	sum2 := 0
	iterator := 0
	found := false
	for scanner.Scan() {
		text := scanner.Text()
		itemMap[iterator] = make(map[rune]bool)
		found = false
		for i, item := range text {
			if itemMap[iterator][item] &&
				i >= len(text)/2 &&
				!found {
				sum1 += priorities[item]
				found = true
			}

			itemMap[iterator][item] = true
		}

		iterator++
		if iterator == 3 {
			for item := range itemMap[0] {
				if itemMap[1][item] && itemMap[2][item] {
					sum2 += priorities[item]
					break
				}
			}
			iterator = 0
		}
	}

	fmt.Println(sum1)
	fmt.Println(sum2)
}
