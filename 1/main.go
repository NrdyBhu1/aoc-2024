package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

func abs(val int) int {
	if val <= 0 {
		return val*(-1)
	}
	return val
}

func arr_count(val int, arr []int) int {
	var count int
	count = 0
	for i := 0; i < len(arr); i++ {
		if arr[i] != val {
			continue
		}
		count++
	}

	return count
}

func quick_sort(arr []int) []int {
	length := len(arr)
	if length == 0 {
		return []int{}
	}
	if length == 1 {
		return arr
	}

	pivot_v := arr[length-1] // Choose the last element as the pivot
	i := 0

	// Partitioning the array
	for k := 0; k < length-1; k++ { // Iterate over the array except the last element
		if arr[k] <= pivot_v {
			arr[i], arr[k] = arr[k], arr[i]
			i++
		}
	}

	// Place the pivot in the correct position
	arr[i], arr[length-1] = arr[length-1], arr[i]
	pivotIndex := i // The index of the pivot after partitioning

	// Create left and right slices
	left := arr[:pivotIndex] // Elements less than or equal to pivot
	right := arr[pivotIndex+1:] // Elements greater than pivot

	// Recursively sort the left and right slices
	left_sorted := quick_sort(left)
	right_sorted := quick_sort(right)

	// Combine the sorted slices
	return append(append(left_sorted, arr[pivotIndex]), right_sorted...)
}

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
	left_sorted = quick_sort(lefts)
	right_sorted = quick_sort(rights)
	
	for i := 0; i < length; i++ {
		distance += abs(left_sorted[i]-right_sorted[i])
	}
	fmt.Println("Part 1: ", distance)

	// part 2
	distance = 0
	for i := 0; i < length; i++ {
		distance += lefts[i] * arr_count(lefts[i], rights)
	}
	fmt.Println("Part 2: ", distance)
}
