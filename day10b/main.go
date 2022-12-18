package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)

	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	calculate(lines)
}

func calculate(lines []string) {
	content := ""
	pixel := 0

	value := 1
	cycle := 0
	cycles := map[int]int{}

	for _, line := range lines {
		split := strings.Split(line, " ")

		if len(split) != 1 {
			param, err := strconv.Atoi(split[1])
			if err != nil {
				panic(err)
			}

			for i := 0; i < 2; i++ {
				cycle += 1
				cycles[cycle] = value

				if isCloseEnough(pixel, value) {
					content += "#"
				} else {
					content += "."
				}
				pixel += 1
				if cycle%40 == 0 {
					content += "\n"
					pixel = 0
				}
			}
			value += param
		} else {
			cycle += 1
			cycles[cycle] = value

			if isCloseEnough(pixel, value) {
				content += "#"
			} else {
				content += "."
			}
			pixel += 1
			if cycle%40 == 0 {
				content += "\n"
				pixel = 0
			}
		}
	}
	fmt.Println(content)
}

func isCloseEnough(a, b int) bool {
	return a == b-1 || a == b || a == b+1
}

func shouldCheck(cycle int) bool {
	return (20-cycle)%40 == 0
}
