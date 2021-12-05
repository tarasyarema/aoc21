package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const INPUT = "in"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func debug(lines []string, v []int) string {
	s := ""

	for _, i := range v {
		s += lines[i] + " "
	}

	return s
}

func main() {
	bs, err := os.ReadFile(INPUT)
	check(err)

	lines := strings.Split(string(bs), "\n")

	n, t := len(lines[0]), len(lines)
	gp := make([]map[int][]int, n)

	for i := 0; i < n; i++ {
		gp[i] = make(map[int][]int)

		gp[i][0] = make([]int, 0)
		gp[i][1] = make([]int, 0)
	}

	for i, line := range lines {
		if len(line) == 0 {
			t -= 1
			continue
		}

		for j, c := range line {
			switch c {
			case '0':
				gp[j][0] = append(gp[j][0], i)
			case '1':
				gp[j][1] = append(gp[j][1], i)
			}
		}
	}

	g, e := 0, 0
	o2, co2 := make([]int, 0), make([]int, 0)

	filter := func(a []int, b []int, eq bool, v int) []int {
		// fmt.Println("filter", eq, v, a, b)

		if len(a) == 0 {
			return b
		} else if len(a) > 1 {
			filtered := make([]int, 0)

			for _, x := range b {
				for _, y := range a {
					if x == y || (eq && x == v) {
						filtered = append(filtered, x)
						break
					}
				}
			}

			return filtered
		}

		return a
	}

	for i := 0; i < len(gp); i++ {
		// fmt.Println(i)
		// fmt.Println(gp[i])
		// fmt.Println()

		if len(gp[i][1]) >= len(gp[i][0]) {
			g += powInt(2, n-i-1)

			if len(o2) == 0 {
				o2 = gp[i][1]
			}

			if len(co2) == 0 {
				co2 = gp[i][0]
			}
		} else {
			e += powInt(2, n-i-1)

			if len(o2) == 0 {
				o2 = gp[i][0]
			}

			if len(co2) == 0 {
				co2 = gp[i][1]
			}
		}

		curr := make(map[int][]int)
		curr[0] = make([]int, 0)
		curr[1] = make([]int, 0)

		if len(o2) > 1 {
			for _, v := range o2 {
				switch lines[v][i] {
				case '0':
					curr[0] = append(curr[0], v)
				case '1':
					curr[1] = append(curr[1], v)
				}
			}
		}

		if i == 0 {
			curr = gp[i]
		}

		// fmt.Println("curr_o2", curr)
		// fmt.Println("pre_o2", o2, debug(lines, o2))

		if len(curr[1]) >= len(curr[0]) {
			o2 = filter(o2, curr[1], len(curr[1]) == len(curr[0]), 1)
		} else {
			o2 = filter(o2, curr[0], false, 0)
		}

		// fmt.Println("post_o2", o2, debug(lines, o2))
		// fmt.Println()

		curr = make(map[int][]int)
		curr[0] = make([]int, 0)
		curr[1] = make([]int, 0)

		if len(co2) > 1 {
			for _, v := range co2 {
				switch lines[v][i] {
				case '0':
					curr[0] = append(curr[0], v)
				case '1':
					curr[1] = append(curr[1], v)
				}
			}
		}

		if i == 0 {
			curr = gp[i]
		}

		// fmt.Println("curr_co2", curr)
		// fmt.Println("pre_co2", co2, debug(lines, co2))

		if len(curr[1]) >= len(curr[0]) {
			co2 = filter(co2, curr[0], len(curr[1]) == len(curr[0]), 0)
		} else {
			co2 = filter(co2, curr[1], false, 0)
		}

		// fmt.Println("post_co2", co2, debug(lines, co2))
		// fmt.Println()
		// fmt.Println()
	}

	fmt.Println(g * e)

	o2Value, err := strconv.ParseInt(lines[o2[0]], 2, 64)
	check(err)

	co2Value, err := strconv.ParseInt(lines[co2[0]], 2, 64)
	check(err)

	// fmt.Println(o2, co2)
	// fmt.Println(lines[o2[0]], lines[co2[0]])
	// fmt.Println(o2Value, co2Value)

	fmt.Println(o2Value * co2Value)
}
