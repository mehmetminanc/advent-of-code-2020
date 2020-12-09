package main

import (
	"bufio"
	"os"
	"regexp"
)

func main() {
	file, _ := os.Open("input/day07.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		exp := regexp.MustCompile("^(.*) bags contain ((\\d*)\\s(.*))*,? bags?.$")
		exp.FindStringSubmatch(text)

	}
}
