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

func main() {
	bs, err := os.ReadFile(INPUT)
	check(err)

	lines := strings.Split(string(bs), "\n")
	ns := make([]int, len(lines))

	sol := 0

	for i, line := range lines {
		if len(line) == 0 {
			continue
		}

		n, err := strconv.Atoi(line)
		check(err)

		ns[i] = n

		if i > 0 && n > ns[i-1] {
			sol += 1
		}
	}

	fmt.Println(sol)
	sol = 0

	var lastPartial *int = nil

	for i := 0; i < len(ns)-2; i++ {
		partial := 0

		for j := 0; j < 3; j++ {
			partial += ns[i+j]
		}

		if lastPartial != nil && partial > *lastPartial {
			sol += 1
		}

		lastPartial = &partial
	}

	fmt.Println(sol)
}
