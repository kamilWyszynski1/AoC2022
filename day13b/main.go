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

type Either struct {
	Value *int     `json:"v,omitempty"`
	List  []Either `json:"l,omitempty"`
}

func (e Either) String() string {
	if e.Value != nil {
		return strconv.Itoa(*e.Value)
	}
	list := make([]string, 0, len(e.List))
	for _, l := range e.List {
		list = append(list, l.String())
	}
	return "[" + strings.Join(list, ",") + "]"
}

func bPtr(b bool) *bool {
	return &b
}

func intPtr(i int) *int {
	return &i
}

func (e Either) order(inx int, another Either) *bool {
	if e.Value != nil && another.Value != nil {
		if *e.Value == *another.Value {
			return nil
		}
		r := *e.Value < *another.Value
		return &r
	}
	if e.Value == nil && another.Value != nil {
		another.List = []Either{{Value: another.Value}}
		another.Value = nil
		return e.order(inx, another)
	}
	if e.Value != nil && another.Value == nil {
		e.List = []Either{{Value: &*e.Value}}
		e.Value = nil
		return e.order(inx, another)
	}
	if e.List != nil && another.List != nil {
		return areInRightOrder(inx, e.List, another.List)
	}
	return bPtr(false)
}

func (e Either) IsValue() bool {
	return e.Value != nil
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func calculate(lines []string) {
	p2 := []Either{{List: []Either{{Value: intPtr(2)}}}}
	p6 := []Either{{List: []Either{{Value: intPtr(6)}}}}
	packets := [][]Either{
		p2, p6,
	}
	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			continue
		}
		first := parse(lines[i])
		i++
		second := parse(lines[i])

		packets = append(packets, first)
		packets = append(packets, second)
	}

	for i := 0; i < len(packets); i++ {
		for j := 0; j < len(packets)-1; j++ {
			ord := areInRightOrder(0, packets[j], packets[j+1])
			if ord == nil {
				continue
			}
			if !*ord {
				packets[j], packets[j+1] = packets[j+1], packets[j]
			}
		}
	}

	inx2 := 0
	inx6 := 0
	for i, packet := range packets {
		if fmt.Sprintf("%s", packet) == fmt.Sprintf("%s", p2) {
			inx2 = i + 1
		}
		if fmt.Sprintf("%s", packet) == fmt.Sprintf("%s", p6) {
			inx6 = i + 1
		}
	}
	fmt.Println(inx2 * inx6)
}

func areInRightOrder(inx int, f, s []Either) *bool {
	for i := 0; i < max(len(f), len(s)); i++ {
		if i >= len(f) {
			return bPtr(true)
		}
		fe := f[i]
		if i >= len(s) {
			return bPtr(false)
		}
		se := s[i]

		o := fe.order(inx, se)
		if o == nil {
			continue
		}
		return o
	}
	return nil
}

// [[4,4],4,4] - []{Either{List: []Either{{Value:4}, {Value: 4}}, Either{Value:4}, Either:{4}}
// [[[]]] - []Either{{List: []Either{{List: []Either{}}}
func parse(line string) []Either {
	if line == "" {
		return nil
	}
	line = line[1 : len(line)-1]
	var eithers []Either
	split := split(line)

	for _, s := range split {
		if !strings.Contains(s, "[") {
			if s != "" {
				// single token
				v, err := strconv.Atoi(s)
				if err != nil {
					panic(err)
				}
				eithers = append(eithers, Either{Value: &v})
			} else {
				eithers = append(eithers, Either{List: []Either{}})
			}
		} else {
			eithers = append(eithers, Either{List: parse(s)})
		}
	}
	return eithers
}

func takeToken(line string) (string, int) {
	add := 0
	for i, ch := range line {
		if ch == ',' && add == 0 {
			return line[1 : i-1], i
		}
		if ch == '[' {
			add += 1
		}
		if ch == ']' {
			add -= 1
		}
	}
	if line[0] == '[' {
		return line[1 : len(line)-1], -1
	}
	return line, -1
}

func split(line string) []string {
	var split []string
	a := 0
	prev := 0
	for i := 0; i < len(line); i++ {
		ch := line[i]
		if ch == '[' {
			a += 1
		}
		if ch == ']' {
			a -= 1
		}
		if ch == ',' && a == 0 {
			split = append(split, line[prev:i])
			prev = i + 1
		}
	}
	split = append(split, line[prev:])
	return split
}
