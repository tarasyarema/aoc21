package main

import (
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

func parse(s string) []int {
	ns := strings.Split(s, ",")

	data := make([]int, 2)

	for i, x := range ns {
		n, err := strconv.Atoi(x)
		check(err)

		data[i] = n
	}

	return data
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func abs(a int) int {
	if a > 0 {
		return a
	}

	return -a
}

func main() {
	t := time.Now()

	defer func() {
		fmt.Printf("Elapsed: %s\n", time.Since(t).String())
	}()

	bs, err := os.ReadFile(INPUT)
	check(err)

	lines := strings.Split(string(bs), "\n")

	N := 1000

	m := make([][]int, N)

	for i := 0; i < N; i++ {
		m[i] = make([]int, N)
	}

	const phase = 2

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		data := strings.Split(line, " -> ")
		l, r := parse(data[0]), parse(data[1])

		xd, yd := r[0]-l[0], r[1]-l[1]
		xm, ym := 1, 1

		d := max(abs(xd), abs(yd))

		if xd != 0 {
			xm = abs(xd)
		}

		if yd != 0 {
			ym = abs(yd)
		}

		unit := []int{
			xd / xm,
			yd / ym,
		}

		if phase == 1 && abs(unit[0]*unit[1]) > 0 {
			continue
		}

		for i := 0; i <= d; i++ {
			m[l[0]+i*unit[0]][l[1]+i*unit[1]] += 1
		}
	}

	// for i := 0; i < N; i++ {
	// 	for j := 0; j < N; j++ {
	// 		y := m[j][i]
	// 		if y == 0 {
	// 			fmt.Printf(".")
	// 		} else {
	// 			fmt.Printf("%d", y)
	// 		}
	// 	}
	// 	fmt.Printf("\n")
	// }

	count := 0

	// Such optimal
	for _, x := range m {
		for _, y := range x {
			if y > 1 {
				count += 1
			}
		}
	}

	fmt.Println(count)

}
