package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("/Users/kamilwyszynski/go/src/stuff/AoC2022/day14b/input.txt")
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

type coords struct {
	x, y int
}

func calculate(lines []string) {
	rocks := map[coords]string{}

	for _, line := range lines {
		split := strings.Split(line, " -> ")

		for i := 0; i < len(split)-1; i++ {
			fromX, fromY := parseCoords(split[i])
			toX, toY := parseCoords(split[i+1])

			if fromX == toX {
				add := 0
				if fromY > toY {
					add = -1
				} else {
					add = 1
				}
				for fromY != toY {
					rocks[coords{fromX, fromY}] = "#"
					fromY += add
				}
				rocks[coords{fromX, fromY}] = "#"
			}
			if fromY == toY {
				add := 0
				if fromX > toX {
					add = -1
				} else {
					add = 1
				}
				for fromX != toX {
					rocks[coords{fromX, fromY}] = "#"

					fromX += add
				}
				rocks[coords{fromX, fromY}] = "#"
			}
		}
	}

	simulate(rocks)
}

func simulate(rocks map[coords]string) {
	// find the floor point
	floor := findLowestPoint(rocks) + 2
	sum := 0

	for {
		// falling
		spawn := coords{500, 0}
		for {
			if spawn.y+1 == floor {
				// sand rests
				rocks[spawn] = "o"
				break
			}
			if _, ok := rocks[coords{spawn.x, spawn.y + 1}]; !ok {
				spawn = coords{spawn.x, spawn.y + 1}
				continue
			}
			// blocked below, what to do
			if _, ok := rocks[coords{spawn.x - 1, spawn.y + 1}]; !ok {
				spawn = coords{spawn.x - 1, spawn.y + 1}
				continue
			}
			if _, ok := rocks[coords{spawn.x + 1, spawn.y + 1}]; !ok {
				spawn = coords{spawn.x + 1, spawn.y + 1}
				continue
			}
			// sand rests
			rocks[spawn] = "o"
			break
		}
		sum += 1
		if spawn.x == 500 && spawn.y == 0 {
			fmt.Println("units ", sum)
			return
		}
	}
}

func findLowestPoint(rocks map[coords]string) int {
	lowest := 0

	for c := range rocks {
		if c.y > lowest {
			lowest = c.y
		}
	}
	return lowest
}

func draw(rocks map[coords]string, floor int) {
	fmt.Print("\033[H\033[2J")

	grid := ""
	for j := 0; j < floor; j++ {
		line := ""
		for i := 485; i <= 600; i++ {
			if j == floor {
				line += "#"
				continue
			}
			if i == 500 && j == 0 {
				line += "+"
				continue
			}
			if v, ok := rocks[coords{i, j}]; ok {
				line += v
			} else {
				line += "."
			}
		}
		grid += line + "\n"
	}
	fmt.Println(grid)
}

func parseCoords(s string) (int, int) {
	split := strings.Split(s, ",")
	if len(split) != 2 {
		panic("split should have len 2")
	}
	x, err := strconv.Atoi(split[0])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(split[1])
	if err != nil {
		panic(err)
	}
	return x, y
}
