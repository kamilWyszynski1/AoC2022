package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sync"
)

func main() {
	f, err := os.Open("/Users/kamilwyszynski/go/src/stuff/AoC2022/day12b/input.txt")
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

type Node struct {
	ID  string
	Val rune
	Str string

	D int

	Children []*Node

	Previous *Node
}

func (n *Node) traverse(fn func(node *Node)) {
	fn(n)
	if n.Previous == nil {
		return
	}
	n.Previous.traverse(fn)
}

func calculate(lines []string) {
	var nodes [][]*Node
	for i, line := range lines {
		n := make([]*Node, 0, len(line))
		for j, ch := range line {
			if ch == 'E' {
				n = append(n, &Node{Val: 'z', Str: string(ch), ID: fmt.Sprintf("%d-%d", i, j)})
			} else {
				n = append(n, &Node{Val: ch, Str: string(ch), ID: fmt.Sprintf("%d-%d", i, j)})
			}
		}
		nodes = append(nodes, n)
	}

	startingPoints := make([]string, 0)
	for i := range nodes {
		for j := range nodes[i] {
			current := nodes[i][j]
			// top
			if i-1 >= 0 {
				n := nodes[i-1][j]
				if n.Val-current.Val < 2 {
					nodes[i][j].Children = append(nodes[i][j].Children, nodes[i-1][j])
				}
			}

			// bottom
			if i+1 < len(nodes) {
				n := nodes[i+1][j]
				if n.Val-current.Val < 2 {
					nodes[i][j].Children = append(nodes[i][j].Children, nodes[i+1][j])
				}
			}

			// left
			if j-1 >= 0 {
				n := nodes[i][j-1]
				if n.Val-current.Val < 2 {
					nodes[i][j].Children = append(nodes[i][j].Children, nodes[i][j-1])
				}
			}
			// right
			if j+1 < len(nodes[i]) {
				n := nodes[i][j+1]
				if n.Val-current.Val < 2 {
					nodes[i][j].Children = append(nodes[i][j].Children, nodes[i][j+1])
				}
			}
			if current.Val == 'a' && len(current.Children) != 0 {
				startingPoints = append(startingPoints, current.ID)
			}
		}
	}
	fmt.Println(len(startingPoints))

	//mutex := &sync.Mutex{}
	wg := &sync.WaitGroup{}
	min := math.MaxInt
	minID := ""

	for _, sp := range startingPoints {
		newNodes := make([][]*Node, len(nodes))
		for i := range nodes {
			l := make([]*Node, len(nodes[i]))
			for j := range nodes[i] {
				cp := *nodes[i][j]
				l[j] = &cp
			}
			newNodes[i] = l
		}

		wg.Add(1)
		go func(sp string, nodes [][]*Node) {
			defer wg.Done()

			var Q []*Node
			var S []*Node

			for i := range nodes {
				for j := range nodes[i] {
					nodes[i][j].Previous = nil
					if nodes[i][j].ID == sp {
						nodes[i][j].D = 0
					} else {
						nodes[i][j].D = math.MaxInt
					}
					Q = append(Q, nodes[i][j])
				}
			}

			for len(Q) != 0 {
				inx := findMin(Q)

				u := Q[inx]
				S = append(S, Q[inx])
				Q = append(Q[:inx], Q[inx+1:]...)

				for _, w := range u.Children {
					if isInQ(w, Q) {
						if w.D > u.D+1 {
							w.D = u.D + 1
							w.Previous = u
						}
					}
				}
			}

			// find start
			var last *Node
			for _, s := range S {
				if s.ID == sp {
					last = s
					break
				}
			}
			sum := 0
			ids := make(map[string]struct{}, 0)
			last.traverse(func(node *Node) {
				sum += 1
				ids[node.ID] = struct{}{}
			})
			sum -= 1
			fmt.Println(sp, sum)
			//mutex.Lock()
			//if sum < min {
			//	minID = sp
			//	min = sum
			//}
			//mutex.Unlock()
		}(sp, newNodes)
	}
	wg.Wait()
	fmt.Println(minID, min)
}

func isInQ(n *Node, Q []*Node) bool {
	for _, q := range Q {
		if n.ID == q.ID {
			return true
		}
	}
	return false
}

func findMin(Q []*Node) int {
	min := math.MaxInt
	minInx := -1

	for i, q := range Q {
		if q.D <= min {
			min = q.D
			minInx = i
		}
	}
	return minInx
}
