package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

const INPUT = "in"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func p(a interface{}) {
	s, err := json.MarshalIndent(a, "", "   ")
	check(err)

	fmt.Println(string(s))
}

func main() {
	t := time.Now()

	defer func() {
		fmt.Printf("Elapsed: %s\n", time.Since(t).String())
	}()

	bs, err := os.ReadFile(INPUT)
	check(err)

	lines := strings.Split(string(bs), "\n")
	m := make([][]int, 0)

	for i, line := range lines {
		if len(line) == 0 {
			continue
		}

		m = append(m, make([]int, len(line)))

		for j, c := range line {
			n, err := strconv.Atoi(string(c))
			check(err)

			m[i][j] = n
		}
	}

	sum := 0
	lows := make([][]int, 0)

	for i, x := range m {
		for j, y := range x {
			ok := true

			if i > 0 && m[i-1][j] <= y {
				ok = false
			}

			if i < len(m)-1 && m[i+1][j] <= y {
				ok = false
			}

			if j > 0 && m[i][j-1] <= y {
				ok = false
			}

			if j < len(x)-1 && m[i][j+1] <= y {
				ok = false
			}

			if ok {
				sum += 1 + y
				lows = append(lows, []int{i, j})
			}
		}
	}

	fmt.Println(sum)

	basins := make([]int, 0)

	for _, low := range lows {
		s := 0

		var curr []int
		nodes := make([][]int, 0)
		nodes = append(nodes, low)

		visited := make(map[int]map[int]bool)

		for i := 0; i < len(lines); i++ {
			visited[i] = map[int]bool{}

			for j := 0; j < len(lines[0]); j++ {
				visited[i][j] = false
			}
		}

		for {
			if len(nodes) == 0 {
				break
			}

			curr, nodes = nodes[0], nodes[1:]

			i, j := curr[0], curr[1]
			y := m[i][j]

			// Mark as visited
			visited[i][j] = true

			if i > 0 && m[i-1][j] > y && m[i-1][j] < 9 && !visited[i-1][j] {
				nodes = append(nodes, []int{i - 1, j})
			}

			if i < len(m)-1 && m[i+1][j] > y && m[i+1][j] < 9 && !visited[i+1][j] {
				nodes = append(nodes, []int{i + 1, j})
			}

			if j > 0 && m[i][j-1] > y && m[i][j-1] < 9 && !visited[i][j-1] {
				nodes = append(nodes, []int{i, j - 1})
			}

			if j < len(m[i])-1 && m[i][j+1] > y && m[i][j+1] < 9 && !visited[i][j+1] {
				nodes = append(nodes, []int{i, j + 1})
			}
		}

		for i := 0; i < len(lines); i++ {
			for j := 0; j < len(lines[0]); j++ {
				w := visited[i][j]

				if w {
					s++
					// fmt.Print("1")
					// } else {
					// fmt.Print("0")
				}
			}

			// fmt.Println()
		}

		basins = append(basins, s)
		// fmt.Println()
	}

	sort.Sort(sort.Reverse(sort.IntSlice(basins)))

	// fmt.Println(basins)
	fmt.Println(basins[0] * basins[1] * basins[2])
}
