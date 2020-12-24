package main

import (
	"bufio"
	mapset "github.com/deckarep/golang-set"
	"os"
	"regexp"
)

type HexCoords struct {
	e  int
	se int
	ne int
}
type XYCoords struct {
	x int
	y int
}

func main() {
	file, _ := os.Open("input/day24.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	coordSet := mapset.NewSet()
	for scanner.Scan() {
		text := scanner.Text()
		coords := XYCoords{
			x: 0,
			y: 0,
		}

		compile := regexp.MustCompile("(e|w|se|sw|ne|nw)")
		submatches := compile.FindAllStringSubmatch(text, 100)

		for _, submatch := range submatches {
			if submatch[0] == "e" {
				coords.x += 2
			} else if submatch[0] == "w" {
				coords.x -= 2
			} else if submatch[0] == "se" {
				coords.x += 1
				coords.y -= 1
			} else if submatch[0] == "nw" {
				coords.x -= 1
				coords.y += 1
			} else if submatch[0] == "ne" {
				coords.x += 1
				coords.y += 1
			} else if submatch[0] == "sw" {
				coords.x -= 1
				coords.y -= 1
			}
		}
		if coordSet.Contains(coords) {
			coordSet.Remove(coords)
		} else {
			coordSet.Add(coords)
		}
	}
	println(len(coordSet.ToSlice()))
}
