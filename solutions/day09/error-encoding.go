package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	part1()
	part2()
}

func part1() {
	file, _ := os.Open("input/day09.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	circular := make([]int64, 25)
	ind := 0
	for scanner.Scan() {
		num, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		if ind >= 25 {
			if !hasMatch(circular, num) {
				println(num)
			}
		}
		circular[ind%25] = num
		ind += 1
	}
}

func hasMatch(circular []int64, target int64) bool {
	for i := 0; i < len(circular)-1; i += 1 {
		for j := i + 1; j < len(circular); j += 1 {
			if circular[i] != circular[j] && circular[i]+circular[j] == target {
				return true
			}
		}
	}
	return false
}

func part2() {
	file, _ := os.Open("input/day09.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	rows := make([]int64, 0)
	for scanner.Scan() {
		num, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		rows = append(rows, num)
	}
	// O(N^3) like a minor boss. There was an algorithm around here for this, somewhere.
	for i := 0; i < len(rows); i += 1 {
		for j := i + 1; j < len(rows); j += 1 {
			sum, max, min := sumPartial(rows, i, j)
			if sum == 373803594 {
				println(max, min, max+min)
			}
		}
	}
}

func sumPartial(rows []int64, i, j int) (int64, int64, int64) {
	sum := int64(0)
	min := int64(^uint(0) >> 1)
	max := int64(-1)
	for k := i; k < j; k += 1 {
		row := rows[k]
		sum += row
		if row > max {
			max = row
		}
		if row < min {
			min = row
		}
	}
	return sum, max, min
}
