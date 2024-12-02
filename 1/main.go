package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
	"aoc/crator/common"
)

func main() {
	var lines []string
	var lefts, rights []int
	var left_sorted, right_sorted []int
	var length, distance int
	fileName := os.Args[1]

	// reading raw data from file
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	}

	// loading data as integer array
	lines = strings.Split(string(data), "\n")
	for i, v := range(lines) {
		line := strings.Fields(v)
		if len(line) != 2 {
			continue
		}
		left_n, err := strconv.Atoi(line[0])
		if err != nil {
			fmt.Println(err, i)
		}
		right_n, err := strconv.Atoi(line[1])
		if err != nil {
			fmt.Println(err, i)
		}
		lefts = append(lefts, left_n)
		rights = append(rights, right_n)
	}

	length = len(lefts)

	// part 1
	distance = 0
	left_sorted = common.QuickSort(lefts)
	right_sorted = common.QuickSort(rights)
	
	for i := 0; i < length; i++ {
		distance += common.Abs(left_sorted[i]-right_sorted[i])
	}
	fmt.Println("Part 1: ", distance)

	// part 2
	distance = 0
	for i := 0; i < length; i++ {
		distance += lefts[i] * common.ArrCount(lefts[i], rights)
	}
	fmt.Println("Part 2: ", distance)
}
