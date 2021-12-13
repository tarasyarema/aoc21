package main

import (
	"encoding/json"
	"fmt"
	"os"
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

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func pm(m [][]int, x, y int) {
	fmt.Println()

	for i := 0; i < min(x, len(m)); i++ {
		for j := 0; j < min(y, len(m[i])); j++ {
			if m[i][j] > 0 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func main() {
	t := time.Now()

	defer func() {
		fmt.Printf("Elapsed: %s\n", time.Since(t).String())
	}()

	input := INPUT

	if len(os.Args) > 1 {
		input = os.Args[1]
	}

	bs, err := os.ReadFile(input)
	check(err)

	lines := strings.Split(string(bs), "\n")

	data := make([][]int, 0)
	folds := make([][2]int, 0)
	pMax := [2]int{0, 0}

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		if strings.Contains(line, "fold") {
			s := strings.Split(line[11:], "=")

			n, err := strconv.Atoi(s[1])
			check(err)

			switch s[0] {
			case "x":
				folds = append(folds, [2]int{0, n})
			case "y":
				folds = append(folds, [2]int{1, n})
			}
		} else {
			s := strings.Split(line, ",")
			point := make([]int, len(s))

			for i, ns := range s {
				n, err := strconv.Atoi(ns)
				check(err)

				point[i] = n
				pMax[i] = max(n, pMax[i])
			}

			data = append(data, point)
		}
	}

	m := make([][]int, pMax[1]+1)

	for i := 0; i <= pMax[1]; i++ {
		m[i] = make([]int, pMax[0]+1)
	}

	for _, d := range data {
		m[d[1]][d[0]] = 1
	}

	x, y := len(m[0]), len(m)
	// pm(m, x, y)

	for i, fold := range folds {
		fmt.Println()
		fmt.Println(fold, x, y, x/2, y/2)

		switch fold[0] {
		case 0:
			for k := 0; k < x; k++ {
				if m[fold[1]][k] > 0 {
					fmt.Println(fold[1], y)
					panic(fmt.Sprintf("X -> %d, %d is %d", k, fold[1], m[k][fold[1]]))
				}
			}

			for k := 0; k < y; k++ {
				for l := 0; l < fold[1]; l++ {
					m[k][l] += m[k][x-1-l]
				}
			}

			x = fold[1]

		case 1:
			for k := 0; k < y; k++ {
				if m[k][fold[1]] > 0 {
					fmt.Println(fold[1], x)
					panic(fmt.Sprintf("Y -> %d, %d is %d", fold[1], k, m[fold[1]][k]))
				}
			}

			for k := 0; k < x; k++ {
				for l := 0; l < fold[1]; l++ {
					m[l][k] += m[y-1-l][k]
				}
			}

			y = fold[1]
		}

		s := 0

		for k := 0; k < x; k++ {
			for l := 0; l < y; l++ {
				if m[l][k] > 0 {
					s++
				}
			}
		}

		fmt.Println(i+1, s)
		// pm(m, y, x)
	}
}
