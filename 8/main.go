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

var letters = map[string]int{
	"f": 9,
	"a": 8,
	"c": 8,
	"d": 7,
	"g": 7,
	"b": 6,
	"e": 4,
}

func newDict() map[string]int {
	return map[string]int{
		"a": 0,
		"b": 0,
		"c": 0,
		"d": 0,
		"e": 0,
		"f": 0,
		"g": 0,
	}
}

func pivotDict(m map[string]int) map[int][]string {
	piv := map[int][]string{
		4: make([]string, 0),
		6: make([]string, 0),
		7: make([]string, 0),
		8: make([]string, 0),
		9: make([]string, 0),
	}

	for k, v := range m {
		piv[v] = append(piv[v], k)
	}

	return piv
}

func parseWord(word string, m map[string]int) int {
	return 0
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func main() {
	t := time.Now()

	defer func() {
		fmt.Printf("Elapsed: %s\n", time.Since(t).String())
	}()

	bs, err := os.ReadFile(INPUT)
	check(err)

	lines := strings.Split(string(bs), "\n")
	orig := pivotDict(letters)
	count, sum := 0, 0

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		parts := strings.Split(line, " | ")

		d := newDict()

		two, three, four := "", "", ""

		for _, word := range strings.Split(parts[0], " ") {
			if len(word) == 2 {
				two = word
			}

			if len(word) == 3 {
				three = word
			}

			if len(word) == 4 {
				four = word
			}

			for _, c := range word {
				d[string(c)] += 1
			}
		}

		mapped := map[string]string{
			"a": "",
			"b": "",
			"c": "",
			"d": "",
			"e": "",
			"f": "",
			"g": "",
		}

		piv := pivotDict(d)

		for k, v := range piv {
			if len(v) > 1 {
				continue
			}

			mapped[v[0]] = orig[k][0]
		}

		// Custom equations
		a := strings.Trim(three, two)
		mapped[a] = "a"

		for _, v := range piv[8] {
			if v != a {
				mapped[v] = "c"
			}
		}

		dm := strings.Trim(four, fmt.Sprintf("%s%s", two, piv[6][0]))
		mapped[dm] = "d"

		for _, v := range piv[7] {
			if v != dm {
				mapped[v] = "g"
			}
		}

		partial := ""

		for _, word := range strings.Split(parts[1], " ") {
			n := ""

			for _, c := range word {
				n += mapped[string(c)]
			}

			n = sortString(n)

			switch n {
			// 0
			case "abcefg":
				partial += "0"
			// 1
			case "cf":
				count += 1
				partial += "1"
				// 2
			case "acdeg":
				partial += "2"
				// 3
			case "acdfg":
				partial += "3"
				// 4
			case "bcdf":
				count += 1
				partial += "4"
				// 5
			case "abdfg":
				partial += "5"
				// 6
			case "abdefg":
				partial += "6"
				// 7
			case "acf":
				count += 1
				partial += "7"
				// 8
			case "abcdefg":
				count += 1
				partial += "8"
				// 9
			case "abcdfg":
				partial += "9"
			}
		}

		m, err := strconv.Atoi(partial)
		check(err)

		sum += m
	}

	fmt.Println(count)
	fmt.Println(sum)
}
