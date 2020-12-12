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
	file, _ := os.Open("input/day12.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	x, y, angle := 0, 0, 0
	for scanner.Scan() {
		ins := scanner.Text()
		val64, _ := strconv.ParseInt(ins[1:], 10, 64)
		val := int(val64)
		switch ins[0] {
		case 'N':
			y += val
		case 'S':
			y -= val
		case 'E':
			x += val
		case 'W':
			x -= val
		case 'L':
			angle += val / 90
		case 'R':
			angle -= val / 90
		case 'F':
			switch angle % 4 {
			case 0:
				x += val
			case 1:
				y += val
			case 2:
				x -= val
			case 3:
				y -= val
			}
		}
	}
	println(x, y)
}

func part2() {
	file, _ := os.Open("input/day12.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	x, y := 0, 0
	wayPointX, wayPointY := 10, 1
	for scanner.Scan() {
		ins := scanner.Text()
		val64, _ := strconv.ParseInt(ins[1:], 10, 64)
		val := int(val64)
		switch ins[0] {
		case 'N':
			wayPointY += val
		case 'S':
			wayPointY -= val
		case 'E':
			wayPointX += val
		case 'W':
			wayPointX -= val
		case 'F':
			x += val * wayPointX
			y += val * wayPointY
		}

		if ins[0] == 'L' || ins[0] == 'R' {
			angle := val / 90
			if ins[0] == 'R' {
				angle = 4 - angle
			}
			switch angle % 4 {
			case 1:
				wayPointX, wayPointY = -wayPointY, wayPointX
			case 2:
				wayPointX, wayPointY = -wayPointX, -wayPointY
			case 3:
				wayPointX, wayPointY = wayPointY, -wayPointX
			}
		}
	}
	println(x, y)
}
