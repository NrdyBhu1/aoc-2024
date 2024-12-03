package main

import (
	"fmt"
	"os"
	"strings"
	"regexp"
	"aoc/crator/common"
)

type Element struct {
	index int
	data string
}

func element_sort(e []Element) []Element {
	length := len(e)
	if length == 0 {
		return []Element{}
	}
	if length == 1 {
		return e
	}

	pivot_e := e[length-1]
	i := 0

	for k := 0; k < length-1; k++ {
		if e[k].index <= pivot_e.index {
			e[i], e[k] = e[k], e[i]
			i++
		}
	}

	e[i], e[length-1] = e[length-1], e[i]
	pivotIndex := i

	left := e[:pivotIndex]
	right := e[pivotIndex+1:]

	left_sorted := element_sort(left)
	right_sorted := element_sort(right)

	return append(append(left_sorted, e[pivotIndex]), right_sorted...)

}

func part2 (expressions []string) {
	var elems []Element
	var res int
	var skip bool
	skip = false
	do_selector := regexp.MustCompile("don't\\(\\)|do\\(\\)")
	mul_selector := regexp.MustCompile("mul\\([0-9]+,[0-9]+\\)")
	num_selector := regexp.MustCompile("[0-9]+,[0-9]+")

	for _, e := range (expressions) {
		do_indices := do_selector.FindAllStringIndex(e, -1)
		do_bytes := do_selector.FindAllString(e, -1)
		elems = []Element{}
		for k := 0; k < len(do_indices); k++ {
			elems = append(elems, Element{ index: do_indices[k][0], data: do_bytes[k] })
		}
		
		mul_indices := mul_selector.FindAllStringIndex(e, -1)
		mul_bytes := mul_selector.FindAllString(e, -1)
		for k := 0; k < len(mul_indices); k++ {
			elems = append(elems, Element{ index: mul_indices[k][0], data: mul_bytes[k] })
		}

		elems = element_sort(elems)
		for _, v := range elems {
			if strings.Contains(v.data, "don't") {
				skip = true
				continue
			}

			if strings.Contains(v.data, "do") {
				skip = false
			}

			if skip {
				continue
			}

			nums := num_selector.FindString(v.data)
			if nums == "" {
				continue
			}
			int_nums := strings.Split(string(nums), ",")
			res += common.Atoi(int_nums[0])*common.Atoi(int_nums[1])
		}
	}
	fmt.Println("Part 2: ", res)
}

func part2red (expressions []string) {
	var res int
	var skip bool
	var dos, donts int
	mul_var_selector := regexp.MustCompile("(don't\\(\\)|do\\(\\))*.mul\\([0-9]+,[0-9]+\\)")
	num_selector := regexp.MustCompile("[0-9]+,[0-9]+")
	for _, e := range (expressions) {
		dos = 0
		donts = 0
		mul_exps := mul_var_selector.FindAllString(e, -1)
		for _, v := range (mul_exps) {
			if strings.Contains(v, "don't") {
				skip = true 
				donts++
				continue
			}

			if strings.Contains(v, "do") && (skip) {
				skip = false
				dos++
			}

			if skip {
				continue
			}

			nums := num_selector.FindString(v)
			int_nums := strings.Split(string(nums), ",")
			res += common.Atoi(int_nums[0])*common.Atoi(int_nums[1])
		}
		fmt.Println(dos, donts)
	}


	fmt.Println("Part 2: ", res)
}

func part1(expressions []string) {
	var res int
	res = 0
	mul_selector := regexp.MustCompile("mul\\([0-9]+,[0-9]+\\)")
	num_selector := regexp.MustCompile("[0-9]+,[0-9]+")
	for _, e := range (expressions) {
		mul_exps := mul_selector.FindAllString(e, -1)
		for _, v := range (mul_exps) {
			nums := num_selector.FindString(v)
			int_nums := strings.Split(string(nums), ",")
			res += common.Atoi(int_nums[0])*common.Atoi(int_nums[1])
		}
	}

	fmt.Println("Part 1: ", res)
}

func main() {
	var lines []string

	fileName := os.Args[1]

	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	}

	lines = strings.Split(string(data), "\n")
	part1(lines)
	part2(lines)
}
