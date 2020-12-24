package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	file, _ := os.Open("input/day13.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	departureTime64, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	scanner.Scan()
	departureTime := int(departureTime64)
	schedulesStr := strings.Split(scanner.Text(), ",")
	schedules := make([]int, 0)

	for i := 0; i < len(schedulesStr); i += 1 {
		scheduleInt, err := strconv.ParseInt(schedulesStr[i], 10, 64)
		if err == nil {
			schedules = append(schedules, int(scheduleInt))
		}
	}
	closestTime := make(map[int]int, 0)

	for i := 0; i < len(schedules); i += 1 {
		for j := 1; j*schedules[i] < departureTime; j += 1 {
			closestTime[schedules[i]] = (j + 1) * schedules[i]
		}
	}
	min := 1_000_000_000
	schedule := -1
	for k, v := range closestTime {
		if v < min {
			min = v
			schedule = k
		}
	}
	println((min - departureTime) * schedule)
}

func part2() {
	file, _ := os.Open("input/day13.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	scanner.Text()
	scanner.Scan()
	schedulesStr := strings.Split(scanner.Text(), ",")
	pairs := make([][2]int, 0)
	for i := 0; i < len(schedulesStr); i += 1 {
		id64, err := strconv.ParseInt(schedulesStr[i], 10, 64)
		if err == nil {
			id := int(id64)
			pairs = append(pairs, [2]int{id, i % id})
		}
	}
	max := max(pairs)
	j := 1
	for true {
		i := j*max[0] + max[1]
		if i < 0 {
			println("overflow")
			break
		}
		if fitsAll(i, pairs) {
			println(i)
			break
		}
		j += 1
	}
}

func fitsAll(num int, pairs [][2]int) bool {
	for i := 0; i < len(pairs); i++ {
		if !fits(num, pairs[i]) {
			return false
		}
	}
	return true
}

func fits(num int, pair [2]int) bool {
	return num%pair[0]-pair[1] == 0
}

func max(pairs [][2]int) [2]int {
	max := [2]int{-1, -1}
	for i := 0; i < len(pairs); i++ {
		if pairs[i][0] > max[0] {
			max = pairs[i]
		}
	}
	return max
}
