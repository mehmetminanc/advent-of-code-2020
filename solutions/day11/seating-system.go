package main

import (
	"bufio"
	"os"
)

func main() {
	file, _ := os.Open("input/day11.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	rows1 := make([][]byte, 0)
	for scanner.Scan() {
		bytes := make([]byte, 100)
		copy(bytes, scanner.Bytes())
		rows1 = append(rows1, bytes)
	}
	rows2 := duplicate(rows1)
	for {
		changeCount := 0
		rows1, changeCount = iteratePart1(rows1)
		if changeCount == 0 {
			println(countOccupied(rows1))
			break
		}
	}
	for {
		changeCount := 0
		rows2, changeCount = iteratePart2(rows2)
		if changeCount == 0 {
			println(countOccupied(rows2))
			break
		}
	}
}

func iteratePart1(rows [][]byte) ([][]byte, int) {
	result := duplicate(rows)
	changeCount := 0
	for y := 0; y < len(rows); y += 1 {
		for x := 0; x < len(rows[0]); x += 1 {
			state := rows[y][x]
			neighbors := countNeighbors(rows, y, x)
			if state == 'L' && neighbors == 0 {
				result[y][x] = '#'
				changeCount += 1
			} else if state == '#' && neighbors >= 4 {
				result[y][x] = 'L'
				changeCount += 1
			}
		}
	}
	return result, changeCount
}

func iteratePart2(rows [][]byte) ([][]byte, int) {
	result := duplicate(rows)
	changeCount := 0
	for y := 0; y < len(rows); y += 1 {
		for x := 0; x < len(rows[0]); x += 1 {
			state := rows[y][x]
			neighbors := countVisible(rows, y, x)
			if state == 'L' && neighbors == 0 {
				result[y][x] = '#'
				changeCount += 1
			} else if state == '#' && neighbors >= 5 {
				result[y][x] = 'L'
				changeCount += 1
			}
		}
	}
	return result, changeCount
}

func countOccupied(rows [][]byte) int {
	count := 0
	for i := 0; i < len(rows); i++ {
		for j := 0; j < len(rows[0]); j++ {
			if rows[i][j] == '#' {
				count += 1
			}
		}
	}
	return count
}

func countVisible(rows [][]byte, r, c int) int {
	return isVisibleOccupied(rows, r, c, -1, -1) +
		isVisibleOccupied(rows, r, c, -1, 0) +
		isVisibleOccupied(rows, r, c, -1, 1) +
		isVisibleOccupied(rows, r, c, 0, -1) +
		isVisibleOccupied(rows, r, c, 0, 1) +
		isVisibleOccupied(rows, r, c, 1, -1) +
		isVisibleOccupied(rows, r, c, 1, 0) +
		isVisibleOccupied(rows, r, c, 1, 1)
}

func isVisibleOccupied(rows [][]byte, r int, c int, y int, x int) int {
	neighborY := r + y
	neighborX := c + x

	if neighborY > -1 && neighborY < len(rows) {
		if neighborX > -1 && neighborX < len(rows[0]) {
			val := rows[neighborY][neighborX]
			if val == '#' {
				return 1
			} else if val == 'L' {
				return 0
			} else if val == '.' {
				return isVisibleOccupied(rows, neighborY, neighborX, y, x)
			}
		}
	}
	return 0
}
func countNeighbors(rows [][]byte, r, c int) int {
	return isOccupied(rows, r, c, -1, -1) +
		isOccupied(rows, r, c, -1, 0) +
		isOccupied(rows, r, c, -1, 1) +
		isOccupied(rows, r, c, 0, -1) +
		isOccupied(rows, r, c, 0, 1) +
		isOccupied(rows, r, c, 1, -1) +
		isOccupied(rows, r, c, 1, 0) +
		isOccupied(rows, r, c, 1, 1)

}

func isOccupied(rows [][]byte, r, c, y, x int) int {
	neighborY := r + y
	neighborX := c + x

	if neighborY > -1 && neighborY < len(rows) {
		if neighborX > -1 && neighborX < len(rows[0]) {
			val := rows[neighborY][neighborX]
			if val == '#' {
				return 1
			}
		}
	}
	return 0
}

func duplicate(rows [][]byte) [][]byte {
	n := len(rows)
	m := len(rows[0])
	duplicate := make([][]byte, n)
	for i := range rows {
		dupRow := make([]byte, m)
		copy(dupRow, rows[i])
		duplicate[i] = dupRow
	}
	return duplicate
}
