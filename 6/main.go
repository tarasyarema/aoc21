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

func main() {
	t := time.Now()

	defer func() {
		fmt.Printf("Elapsed: %s\n", time.Since(t).String())
	}()

	bs, err := os.ReadFile(INPUT)
	check(err)

	numString := strings.Split(strings.Split(string(bs), "\n")[0], ",")
	data := make([]int, len(numString))
	data2 := make(map[int]int)

	for i := 0; i < 9; i++ {
		data2[i] = 0
	}

	for i, s := range numString {
		n, err := strconv.Atoi(s)
		check(err)

		data[i] = n
		data2[n] += 1
	}

	for i := 0; i < 80; i++ {
		newFish := 0

		for j, x := range data {
			if x == 0 {
				newFish += 1

				data[j] = 6
			} else {
				data[j] = x - 1
			}
		}

		for j := 0; j < newFish; j++ {
			data = append(data, 8)
		}
	}

	fmt.Println(len(data))

	for i := 0; i < 256; i++ {
		newOnes := data2[0]

		for j := 1; j < 9; j++ {
			data2[j-1] = data2[j]
		}

		data2[8] = newOnes
		data2[6] += newOnes
	}

	sum := 0

	for i := 0; i < 9; i++ {
		sum += data2[i]
	}

	fmt.Println(sum)
}
