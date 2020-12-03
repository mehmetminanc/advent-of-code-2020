package main

import (
	"bufio"
	"os"
)

func main() {
	file, _ := os.Open("input/day03.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	rows := make([]string, 0)

	for scanner.Scan() {
		rows = append(rows, scanner.Text())
	}

	result := countPerSlope(rows, 1, 1) *
		countPerSlope(rows, 1, 3) *
		countPerSlope(rows, 1, 5) *
		countPerSlope(rows, 1, 7) *
		countPerSlope(rows, 2, 1)

	println(result)
}

func countPerSlope(rows []string, down int, right int) int {

	rc, cc := len(rows), len(rows[0])

	trees, x := 0, 0
	for y := 0; y < rc; y += down {
		if rows[y][x%cc] == '#' {
			trees += 1
		}
		x += right
	}

	return trees
}
