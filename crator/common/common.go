package common

import (
	"fmt"
	"strconv"
)

func Atoi(num string) int {
	data, err := strconv.Atoi(num)
	if err != nil {
		fmt.Println(err)
	}

	return data
}

func Abs(val int) int {
	if val <= 0 {
		return val*(-1)
	}
	return val
}

func ArrCount(val int, arr []int) int {
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


func QuickSort(arr []int) []int {
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
	left_sorted := QuickSort(left)
	right_sorted := QuickSort(right)

	// Combine the sorted slices
	return append(append(left_sorted, arr[pivotIndex]), right_sorted...)
}
