package main

import (
	"fmt"
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

func checkBingo(a []int, b []int) bool {
	// Check horizontal
	for i := 0; i < 5; i++ {
		won := true

		for j := 0; j < 5; j++ {
			if b[i*5+j] != -1 {
				won = false
			}
		}

		if won {
			return true
		}
	}

	// Check vertical
	for i := 0; i < 5; i++ {
		won := true

		for j := 0; j < 5; j++ {
			if b[i+j*5] != -1 {
				won = false
			}
		}

		if won {
			return true
		}
	}

	return false
}

func solveBingo(draws []int, b []int) (int, int) {
	curr := make([]int, len(b))
	copy(curr, b)

	for i, d := range draws {
		for j := range curr {
			if curr[j] == d {
				curr[j] = -1
			}
		}

		if checkBingo(b, curr) {
			// fmt.Println(draws[:i+1])
			// for i := 0; i < 5; i++ {
			// 	fmt.Println(b[i*5 : i*5+5])
			// }
			// fmt.Println()
			// for i := 0; i < 5; i++ {
			// 	fmt.Println(curr[i*5 : i*5+5])
			// }

			sum := 0

			for i, v := range curr {
				if v != -1 {
					sum += b[i]
				}
			}

			// fmt.Println(i, sum, d)
			// fmt.Println()

			return i, sum * d
		}
	}

	return -1, 0
}

func main() {
	bs, err := os.ReadFile(INPUT)
	check(err)

	lines := strings.Split(string(bs), "\n")

	drawsStr := strings.Split(lines[0], ",")
	draws := make([]int, len(drawsStr))

	for i, s := range drawsStr {
		n, err := strconv.Atoi(s)
		check(err)

		draws[i] = n
	}

	bingo := make([][]int, 0)
	nn := 0

	for i := 2; i < len(lines); i += 6 {
		bingo = append(bingo, make([]int, 0))

		for _, line := range lines[i : i+5] {
			for _, c := range strings.Split(line, " ") {
				if c == "" {
					continue
				}

				n, err := strconv.Atoi(c)
				check(err)

				bingo[nn] = append(bingo[nn], n)
			}
		}

		if len(bingo[nn]) != 25 {
			panic("Wrong len")
		}

		nn += 1
	}

	winner := []int{len(draws) + 1, 0}
	winnerLast := []int{0, 0}

	for _, b := range bingo {
		pos, value := solveBingo(draws, b)

		if pos > -1 && pos < winner[0] {
			winner[0] = pos
			winner[1] = value
		}

		if pos > -1 && pos > winnerLast[0] {
			winnerLast[0] = pos
			winnerLast[1] = value
		}
	}

	fmt.Println(winner)
	fmt.Println(winnerLast)
}
