package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type node struct {
	children map[string]node
	fileSize int
}

func main() {
	f, err := os.Open("day7/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	lines := make(chan string)
	scanner.Scan()
	n := node{children: make(map[string]node)}
	go func() {
		for scanner.Scan() {
			text := scanner.Text()
			lines <- text
		}
		close(lines)
	}()
	n.eval(lines)
	fmt.Println(n.sum(100000))
	fmt.Println(n.deleteDirectory(30000000, 70000000))
}

func (n *node) eval(lines <-chan string) {
	for line := range lines {
		if line == "$ cd .." {
			return
		}

		splitLine := strings.Split(line, " ")

		if splitLine[0] == "$" {
			if splitLine[1] == "cd" {
				child := n.getChild(splitLine[2])
				child.eval(lines)
			}
		} else if splitLine[0] == "dir" {
			_, ok := n.children[splitLine[1]]
			if !ok {
				n.children[splitLine[1]] = node{
					children: make(map[string]node),
				}
			}
		} else {
			amount, err := strconv.Atoi(splitLine[0])
			if err != nil {
				panic(err)
			}
			n.children[splitLine[1]] = node{fileSize: amount}
		}
	}
}

func (n *node) getChild(name string) node {
	splitPath := strings.Split(name, "/")
	child, ok := n.children[splitPath[0]]
	if !ok {
		child = node{
			children: make(map[string]node),
		}
		n.children[splitPath[0]] = child
	}

	if len(splitPath) > 1 {
		return child.getChild(strings.Join(splitPath, "/"))
	}

	return child
}

func (n *node) size() int {
	sum := 0
	for _, child := range n.children {
		sum += child.size()
	}
	sum += n.fileSize
	return sum
}

func (n node) flatten(parent string) map[string]node {
	output := make(map[string]node)
	if len(n.children) == 0 {
		return output
	}

	output[parent] = n
	for name, child := range n.children {
		for childName, flatChild := range child.flatten(parent + "/" + name) {
			output[childName] = flatChild
		}
	}
	return output
}

func (n *node) sum(threshhold int) int {
	sum := 0

	for _, child := range n.flatten("") {
		size := child.size()
		//fmt.Println(size, childName, ":", fmt.Sprintf("%+v", child))
		if size < threshhold {
			sum += size
		}
	}

	return sum
}

func (n *node) deleteDirectory(threshhold int, totalSpace int) int {
	freeSpace := totalSpace - n.size()
	spaceNeeded := threshhold - freeSpace
	toBeDeleted := totalSpace
	for _, child := range n.flatten("") {
		size := child.size()
		if size >= spaceNeeded && size < toBeDeleted {
			toBeDeleted = size
		}
	}

	return toBeDeleted
}
