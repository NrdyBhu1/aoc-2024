package main

import (
	"fmt"
	"os"
	"strings"
	"aoc/crator/common"
)

func check_safety (levels []int) bool {
	inc := true
	dec := true

	for i := 0; i < len(levels)-1; i++ {
		diff := common.Abs(levels[i+1] - levels[i])

		if !(diff >= 1 && diff <= 3) {
			return false
		}

		if levels[i] < levels[i+1] {
            dec = false
        } else if levels[i] > levels[i+1] {
            inc = false
        }
	}

	return inc || dec
}

func check_safety2 (levels []int) bool {
	for i := 0; i < len(levels); i++ {
		temp_level := append([]int{}, levels[:i]...)
		temp_level = append(temp_level, levels[i+1:]...)

		if check_safety(temp_level) {
			return true
		}
	}

	return false
}


func part2(reports [][]int) {
	var safeReports int
	safeReports = 0
	for _,report := range(reports) {
		if check_safety(report) || check_safety2(report) {
			safeReports++
		}
	}

	fmt.Println("Part 2: ", safeReports)
}

func part1(reports [][]int) {
	var safeReports int
	safeReports = 0
	for j := 0; j < len(reports); j++ {
		if check_safety(reports[j]) {
			safeReports++
		}
	}
	fmt.Println("Part 1: ", safeReports)
}

func main() {
	var lines []string
	var reports [][]int
	
	fileName := os.Args[1]

	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	}

	lines = strings.Split(string(data), "\n")
	lines = lines[:len(lines)-1]
	for i, v := range(lines) {
		levels := strings.Fields(v)
		reports = append(reports, []int{})
		for _, lvl := range(levels) {
			reports[i] = append(reports[i], common.Atoi(lvl))
		}
	}
	part1(reports)
	part2(reports)
}
