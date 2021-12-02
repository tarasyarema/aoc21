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

	h, d, a := 0, 0, 0

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		line_str := strings.Split(line, " ")

		n, err := strconv.Atoi(line_str[1])
		check(err)

		switch line_str[0] {
		case "forward":
			h += n
			d += (a * n)
		case "down":
			a += n
		case "up":
			a -= n
		}
	}

	fmt.Println(h * d)
}
