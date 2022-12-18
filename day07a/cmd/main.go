package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type nodeType uint8

const (
	dir nodeType = iota
	file
)

type node struct {
	name string

	parent   *node
	children []*node

	tpe  nodeType
	size int
}

func (n node) find(fn func(n node)) {
	fn(n)
	for _, ch := range n.children {
		ch.find(fn)
	}
}

func (n *node) calculate() {
	if n.tpe == file {
		return
	}
	sum := 0
	for _, ch := range n.children {
		ch.calculate()
		sum += ch.size
	}
	n.size = sum
}

func (n node) Print() {
	n.print(0)
}

func (n node) print(indentation int) {
	prefix := ""
	for i := 0; i < indentation; i++ {
		prefix = "\t" + prefix
	}
	switch n.tpe {
	case dir:
		fmt.Printf("%s- %s (dir)\n", prefix, n.name)
	case file:
		fmt.Printf("%s- %s (file, size=%d)\n", prefix, n.name, n.size)
	}
	for _, ch := range n.children {
		ch.print(indentation + 1)
	}
}

func (n node) PrintSize() {
	n.printSize(0)
}

func (n node) printSize(indentation int) {
	prefix := ""
	for i := 0; i < indentation; i++ {
		prefix = "\t" + prefix
	}
	switch n.tpe {
	case dir:
		fmt.Printf("%s- %s (dir, size=%d)\n", prefix, n.name, n.size)
	case file:
		fmt.Printf("%s- %s (file, size=%d)\n", prefix, n.name, n.size)
	}
	for _, ch := range n.children {
		ch.printSize(indentation + 1)
	}
}

func fileToSlice() []string {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}

func calculate(lines []string) {
	var root *node = nil

	for _, line := range lines {
		split := strings.Split(line, " ")

		switch split[0] {
		case "$":
			switch split[1] {
			case "cd":
				switch name := split[2]; name {
				case "..":
					root = root.parent
				default:
					if root == nil {
						root = &node{name: split[2], tpe: dir}
					} else {
						new := &node{name: split[2], tpe: dir, parent: root}
						root.children = append(root.children, new)
						root = new
					}
				}
			}
		case "dir":
			// skip for now
		default:
			size, err := strconv.Atoi(split[0])
			if err != nil {
				log.Panicf("could not parse %s size", split[0])
			}
			root.children = append(
				root.children,
				&node{
					name:   split[1],
					tpe:    file,
					parent: root,
					size:   size,
				},
			)
		}
	}
	// go back to very begining
	for root.parent != nil {
		root = root.parent
	}
	root.calculate()
	root.PrintSize()

	// ### Part one ###
	// result := 0

	// root.find(func(n node) {
	// 	if n.tpe == dir && n.size < 100000 {
	// 		result += n.size
	// 	}
	// })
	// fmt.Println(result)

	// ### Part two ###
	missing := 30000000 - (70000000 - root.size)
	minimal := math.MaxInt
	root.find(func(n node) {
		if n.tpe == dir && n.size > missing && n.size < minimal {
			minimal = n.size
		}
	})
	fmt.Println(minimal)
}

func main() {
	calculate(fileToSlice())
}
