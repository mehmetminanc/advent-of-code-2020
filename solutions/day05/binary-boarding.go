package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("input/day05.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var part2visual [127][8]bool

	max := 0
	for scanner.Scan() {
		row, column := findRowColumn(scanner.Text())
		part2visual[row][column] = true
		seatId := calculateSeatId(row, column)
		if seatId > max {
			max = seatId
		}
	}
	println(max)
	printVisual(part2visual)
}

func printVisual(visual [127][8]bool) {
	fmt.Printf("    0 1 2 3 4 5 6 7\n")
	for indR, rows := range visual {
		fmt.Printf("%-4d", indR)
		for _, c := range rows {
			if c {
				print("# ")
			} else {
				print(". ")
			}
		}
		print("\n")
	}
}

func findRowColumn(ticket string) (r int, c int) {
	rowInd := 63
	rowNext := 32
	for i := 0; i < 7; i++ {
		chr := ticket[i]
		if chr == 'F' {
			rowInd -= rowNext
		} else {
			rowInd += max(rowNext, 1)
		}
		rowNext = rowNext / 2
	}
	colInd := 3
	colNext := 2
	for i := 7; i < 10; i++ {
		chr := ticket[i]
		if chr == 'L' {
			colInd -= colNext
		} else {
			colInd += max(colNext, 1)
		}
		colNext = colNext / 2
	}
	return rowInd, colInd
}

func calculateSeatId(r int, c int) int {
	return 8*r + c
}

func max(a int, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
