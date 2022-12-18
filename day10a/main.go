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
			}
			value += param
		} else {
			cycle += 1
			cycles[cycle] = value
		}
	}

	sum := 0
	for k, v := range cycles {
		if shouldCheck(k) {
			sum += k * v
		}
	}

	fmt.Println(sum)
}

func shouldCheck(cycle int) bool {
	return (20-cycle)%40 == 0
}
