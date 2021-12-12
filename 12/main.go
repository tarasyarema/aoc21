package main

import (
	"encoding/json"
	"fmt"
	"os"
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

func isBig(s string) bool {
	for _, c := range s {
		// Assume all will
		if c < 97 {
			return true
		}
	}

	return false
}

func solve(g map[string][]string, vis map[string]bool, curr string, p int) int {
	vis[curr] = true

	if curr == "end" {
		return 1
	}

	branch := p

	for _, n := range g[curr] {
		if !isBig(n) && vis[n] {
			continue
		}

		visCopy := map[string]bool{}

		for k, v := range vis {
			visCopy[k] = v
		}

		branch += solve(g, visCopy, n, p)
	}

	return branch
}

func solve2(g map[string][]string, vis map[string]int, curr string, p int) int {
	vis[curr] += 1

	if curr == "end" {
		return 1
	}

	branch := p

	for _, n := range g[curr] {
		// Avoid loops to the start
		if n == "start" {
			continue
		}

		done := false
		visCopy := map[string]int{}

		for k, v := range vis {
			if !isBig(k) && k != "end" && k != "start" && v == 2 {
				done = true
			}

			visCopy[k] = v
		}

		if !isBig(n) && vis[n] > 0 && done {
			continue
		}

		branch += solve2(g, visCopy, n, p)
	}

	return branch
}

func main() {
	t := time.Now()

	defer func() {
		fmt.Printf("Elapsed: %s\n", time.Since(t).String())
	}()

	bs, err := os.ReadFile(INPUT)
	check(err)

	lines := strings.Split(string(bs), "\n")
	g := map[string][]string{}

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		nodes := strings.Split(line, "-")

		for _, node := range nodes {
			if _, ok := g[node]; !ok {
				g[node] = []string{}
			}
		}

		g[nodes[0]] = append(g[nodes[0]], nodes[1])
		g[nodes[1]] = append(g[nodes[1]], nodes[0])
	}

	fmt.Println(solve(g, map[string]bool{}, "start", 0))
	fmt.Println(solve2(g, map[string]int{}, "start", 0))
}
