package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Password struct {
	min      int64
	max      int64
	letter   string
	password string
}

func isValidPart1(p Password) bool {
	count := int64(0)
	for _, c := range p.password {
		if string(c) == p.letter {
			count += 1
		}
	}
	return count >= p.min && count <= p.max
}

func isValidPart2(p Password) bool {
	return (p.password[p.min-1] == p.letter[0]) != (p.password[p.max-1] == p.letter[0])
}

func main() {
	println(counter(isValidPart1))
	println(counter(isValidPart2))
}

func counter(valid func(password Password) bool) int {
	file, _ := os.Open("input/day02.txt")
	defer file.Close()
	count := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		parts := strings.Split(text, ":")
		partsLeft := strings.Split(parts[0], " ")
		partsMinMax := strings.Split(partsLeft[0], "-")

		p := Password{}
		p.min, _ = strconv.ParseInt(partsMinMax[0], 10, 64)
		p.max, _ = strconv.ParseInt(partsMinMax[1], 10, 64)
		p.password = strings.TrimSpace(parts[1])
		p.letter = strings.TrimSpace(partsLeft[1])
		if valid(p) {
			count += 1
		}
	}
	return count
}
