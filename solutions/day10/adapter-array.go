package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

func main() {
	rows := readRows()
	part1(rows)
	part2(rows)
}

func readRows() []int {
	file, _ := os.Open("input/day10.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	rows := make([]int, 0)
	rows = append(rows, 0)
	for scanner.Scan() {
		val64, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		rows = append(rows, int(val64))
	}
	sort.Ints(rows)
	rows = append(rows, rows[len(rows)-1]+3)
	return rows
}

func part1(rows []int) {
	histogram := map[int]int{
		1: 0,
		2: 0,
		3: 0,
	}
	for i := 0; i < len(rows)-1; i += 1 {
		histogram[rows[i+1]-rows[i]] += 1
	}
	println(histogram[1] * histogram[3])
}

func part2(rows []int) {
	counts := make([]int, len(rows))
	counts[len(counts)-1] = 0
	counts[len(counts)-2] = 1
	for i := len(rows) - 1; i >= 0; i -= 1 {
		for j := 1; j <= 3 && i-j >= 0; j += 1 {
			if rows[i]-rows[i-j] <= 3 {
				counts[i-j] += counts[i]
			}
		}
	}
	println(counts[0])
}
