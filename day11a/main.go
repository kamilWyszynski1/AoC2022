package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("/Users/kamilwyszynski/go/src/stuff/AoC2022/day11a/input.txt")
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

type Monkey struct {
	ID        int
	Items     []int
	Operation func(value int) int
	Test      func(monkey *Monkey)

	Inspections int
}

func (m *Monkey) String() string {
	return fmt.Sprintf("Monkey %d inspected items %d times\n", m.ID, m.Inspections)
}

func (m *Monkey) Process() {
	for len(m.Items) != 0 {
		m.Inspections += 1

		m.Items[0] = m.Operation(m.Items[0])
		//m.Items[0] /= 3
		m.Test(m)
	}
}

func parseMonkeys(lines []string) []*Monkey {
	monkeys := make(map[int]*Monkey)

	id := -1
	var current *Monkey = nil
	for i := 0; i < len(lines); i += 1 {
		line := lines[i]

		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}

		switch true {
		case strings.HasPrefix(line, "Monkey"):
			if current != nil {
				monkeys[id] = current
				id += 1
			} else {
				id = 0
			}
			current = &Monkey{ID: id}
		case strings.HasPrefix(line, "Starting items"):
			items := strings.Split(strings.Split(line, ": ")[1], ",")

			parsed := make([]int, 0, len(items))
			for _, item := range items {
				value, err := strconv.Atoi(strings.TrimSpace(item))
				if err != nil {
					panic(err)
				}
				parsed = append(parsed, value)
			}
			current.Items = parsed

		case strings.HasPrefix(line, "Operation"):
			split := strings.Split(line, " ")

			valueStr := split[len(split)-1]
			var (
				value int
				err   error
			)
			if valueStr != "old" {
				value, err = strconv.Atoi(valueStr)
				if err != nil {
					panic(err)
				}

			}
			switch token := split[len(split)-2]; token {
			case "+":
				if valueStr == "old" {
					current.Operation = func(o int) int {
						return o + o
					}
				} else {
					current.Operation = func(o int) int {
						return o + value
					}
				}
			case "*":
				if valueStr == "old" {
					current.Operation = func(o int) int {
						return o * o
					}
				} else {
					current.Operation = func(o int) int {
						return o * value
					}
				}
			default:
				panic(fmt.Sprintf("unexpected token: %s", token))
			}
		case strings.HasPrefix(line, "Test"):
			split := strings.Split(line, " ")
			div, err := strconv.Atoi(split[len(split)-1])
			if err != nil {
				panic(err)
			}

			// if true
			i += 1
			line = lines[i]
			split = strings.Split(line, " ")
			trueMonkeyID, err := strconv.Atoi(split[len(split)-1])
			if err != nil {
				panic(err)
			}
			// if false
			i += 1
			line = lines[i]
			split = strings.Split(line, " ")
			falseMonkeyID, err := strconv.Atoi(split[len(split)-1])
			if err != nil {
				panic(err)
			}

			current.Test = func(monkey *Monkey) {
				value := monkey.Items[0]
				fmt.Printf("testing %d item\n", value)

				if value%div == 0 {
					item := monkey.Items[0]
					monkey.Items = monkey.Items[1:]

					m, ok := monkeys[trueMonkeyID]
					if !ok {
						panic(fmt.Sprintf("monkey %d tried to throw item %d to %d", id, item, trueMonkeyID))
					}
					m.Items = append(m.Items, item)
				} else {
					item := monkey.Items[0]
					monkey.Items = monkey.Items[1:]

					m, ok := monkeys[falseMonkeyID]
					if !ok {
						panic(fmt.Sprintf("monkey %d tried to throw item %d to %d", id, item, trueMonkeyID))
					}
					m.Items = append(m.Items, item)
				}
			}
		default:
			panic(fmt.Sprintf("default on %s", line))
		}
	}
	monkeys[id] = current

	slice := make([]*Monkey, 0, len(monkeys))
	for _, v := range monkeys {
		slice = append(slice, v)
	}

	sort.Slice(slice, func(i, j int) bool {
		return slice[i].ID < slice[j].ID
	})
	return slice
}

func calculate(lines []string) {
	monkeys := parseMonkeys(lines)

	for i := 1; i <= 10000; i++ {
		for _, monkey := range monkeys {
			monkey.Process()
		}

		for _, round := range []int{1, 20, 1000, 2000, 3000} {
			if i == round {
				fmt.Println("==", "After round", i, "==")
				fmt.Println(monkeys)
				fmt.Println()
			}
		}
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].Inspections < monkeys[j].Inspections
	})

	fmt.Println(monkeys[len(monkeys)-1].Inspections * monkeys[len(monkeys)-2].Inspections)
}
