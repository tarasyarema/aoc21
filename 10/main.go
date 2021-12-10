package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
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
	sum := 0

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		var x rune
		a := []rune(string(line[0]))
		s := 0

		for _, c := range line[1:] {

			switch c {
			case '(':
				a = append(a, c)
			case '[':
				a = append(a, c)
			case '{':
				a = append(a, c)
			case '<':
				a = append(a, c)

			case ')':
				x, a = a[len(a)-1], a[:len(a)-1]

				if x != '(' {
					s += 3
					break
				}
			case ']':
				x, a = a[len(a)-1], a[:len(a)-1]

				if x != '[' {
					s += 57
					break
				}
			case '}':
				x, a = a[len(a)-1], a[:len(a)-1]

				if x != '{' {
					s += 1197
					break
				}
			case '>':
				x, a = a[len(a)-1], a[:len(a)-1]

				if x != '<' {
					s += 25137
					break
				}
			}
		}

		sum += s
	}

	fmt.Println(sum)

	scores := make([]int, 0)

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		var x rune
		a := []rune(string(line[0]))
		bad := false

		for _, c := range line[1:] {
			switch c {
			case '(':
				a = append(a, c)
			case '[':
				a = append(a, c)
			case '{':
				a = append(a, c)
			case '<':
				a = append(a, c)

			case ')':
				x, a = a[len(a)-1], a[:len(a)-1]

				if x != '(' {
					bad = true
					break
				}
			case ']':
				x, a = a[len(a)-1], a[:len(a)-1]

				if x != '[' {
					bad = true
					break
				}
			case '}':
				x, a = a[len(a)-1], a[:len(a)-1]

				if x != '{' {
					bad = true
					break
				}
			case '>':
				x, a = a[len(a)-1], a[:len(a)-1]

				if x != '<' {
					bad = true
					break
				}
			}
		}

		if bad {
			continue
		}

		s := 0

		for {
			if len(a) == 0 {
				break
			}

			s *= 5
			x, a = a[len(a)-1], a[:len(a)-1]

			switch x {
			case '(':
				s += 1
			case '[':
				s += 2
			case '{':
				s += 3
			case '<':
				s += 4
			}
		}

		scores = append(scores, s)
	}

	sort.Ints(scores)
	fmt.Println(scores[len(scores)/2])
}
