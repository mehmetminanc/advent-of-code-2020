package main

import (
	"bufio"
	"bytes"
	"os"
	"strings"
)

func main() {
	part1()
	part2()
}

func part2() {
	file, _ := os.Open("input/day06.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(ScanDoubleLines)

	sum := 0
	for scanner.Scan() {
		text := scanner.Text()
		passengers := strings.Split(text, "\n")
		choices := make([][26]bool, 0)
		for _, passenger := range passengers {
			trimmed := strings.Trim(passenger, "\n")
			if len(trimmed) != 0 {
				choice := [26]bool{false}
				fillArr(&choice, trimmed)
				choices = append(choices, choice)
			}
		}
		sum += countChoices(choices)
	}
	println(sum)
}

func countChoices(choices [][26]bool) int {
	count := 0
	for i := 'a'; i <= 'z'; i++ {
		all := true
		for _, choice := range choices {
			if !choice[i-'a'] {
				all = false
				break
			}
		}
		if all {
			count += 1
		}
	}
	return count
}

func part1() {
	file, _ := os.Open("input/day06.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(ScanDoubleLines)

	arr := [26]bool{false}
	sumPart1 := 0
	for scanner.Scan() {
		text := scanner.Text()
		emptyArr(&arr)
		fillArr(&arr, text)
		sumPart1 += countArr(&arr)
	}
	println(sumPart1)
}

func fillArr(arr *[26]bool, text string) {
	for _, c := range text {
		if c >= 'a' && c <= 'z' {
			arr[c-'a'] = true
		}
	}
}

func emptyArr(arr *[26]bool) {
	for i := 0; i < len(arr); i++ {
		arr[i] = false
	}
}

func countArr(arr *[26]bool) int {
	count := 0

	for i := 0; i < len(arr); i++ {
		if arr[i] {
			count += 1
		}
	}
	return count
}

func dropCR(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '\r' {
		return data[0 : len(data)-1]
	}
	return data
}

func ScanDoubleLines(data []byte, atEOF bool) (advance int, token []byte, err error) {
	match := []byte{'\n', '\n'}
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.Index(data, match); i >= 0 {
		// We have a full newline-terminated line.
		return i + 2, dropCR(data[0 : i+2]), nil
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), dropCR(data), nil
	}
	// Request more data.
	return 0, nil, nil
}
