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

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

func main() {
	t := time.Now()

	defer func() {
		fmt.Printf("Elapsed: %s\n", time.Since(t).String())
	}()

	bs, err := os.ReadFile(INPUT)
	check(err)

	numString := strings.Split(strings.Split(string(bs), "\n")[0], ",")

	data := make([]int, len(numString))
	ma, mi := 0, 100000

	for i, s := range numString {
		n, err := strconv.Atoi(s)
		check(err)

		data[i] = n

		if n > ma {
			ma = n
		}

		if n < mi {
			mi = n
		}
	}

	best := 10000000
	best2 := 10000000000

	for i := mi; i <= ma; i++ {
		sum := 0
		sum2 := 0

		for _, x := range data {
			sum += abs(x - i)

			tmp := abs(x - i)
			sum2 += (tmp * (tmp + 1)) / 2
		}

		if sum < best {
			best = sum
		}

		if sum2 < best2 {
			best2 = sum2
		}
	}

	fmt.Println(best)
	fmt.Println(best2)
}
