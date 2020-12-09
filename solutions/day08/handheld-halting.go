package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	operation int
	value     int
	visited   bool
}

func main() {
	file, _ := os.Open("input/day08.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	executable := make([]Instruction, 0)
	jumps := make([]int, 0)
	i := 0
	for scanner.Scan() {
		text := scanner.Text()
		split := strings.Split(text, " ")
		op := split[0]
		val64, _ := strconv.ParseInt(split[1], 10, 64)
		val := int(val64)

		instruction := Instruction{}
		instruction.value = val
		instruction.visited = false
		if op == "nop" {
			instruction.operation = 0
		} else if op == "acc" {
			instruction.operation = 1
		} else {
			instruction.operation = 2
			jumps = append(jumps, i)
		}
		executable = append(executable, instruction)
		i += 1
	}
	executablePart2 := make([]Instruction, len(executable))
	copy(executablePart2, executable)

	part1(executable)
	part2(executablePart2, jumps)
}

func part1(executable []Instruction) {
	pc := 0
	acc := 0

	for true {
		instruction := &executable[pc]
		if instruction.visited {
			fmt.Printf("Part1: %d \n", acc)
			break
		} else {
			instruction.visited = true
			if instruction.operation == 0 {
				pc += 1
			} else if instruction.operation == 1 {
				acc += instruction.value
				pc += 1
			} else {
				pc += instruction.value
			}
		}
	}
}

func part2(executable []Instruction, jumps []int) {
	for _, jump := range jumps {
		executableCopy := make([]Instruction, len(executable))
		copy(executableCopy, executable)

		pc := 0
		acc := 0
		for true {
			if pc >= len(executableCopy) {
				fmt.Printf("Part2: %d \n", acc)
				break
			}

			instruction := &executableCopy[pc]
			if pc > len(executableCopy) {
				fmt.Printf("Part2: %d \n", acc)
			}
			if instruction.visited {
				break
			} else {
				instruction.visited = true
				if pc == jump || instruction.operation == 0 {
					pc += 1
				} else if instruction.operation == 1 {
					acc += instruction.value
					pc += 1
				} else {
					jumps = append(jumps, pc)
					pc += instruction.value
				}
			}
		}
	}
}
