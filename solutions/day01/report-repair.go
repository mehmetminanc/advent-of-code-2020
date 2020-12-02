package main

import (
	"bufio"
	mapset "github.com/deckarep/golang-set"
	"os"
	"strconv"
)

func reportRepairPart1(year int64, set mapset.Set) int64 {
	iterator := set.Iterator()
	for elem := range iterator.C {
		var remainder = year - elem.(int64)
		if set.Contains(remainder) {
			return remainder * elem.(int64)
		}
	}
	return -1
}

func reportRepairPart2(year int64, set mapset.Set) int64 {
	product := set.CartesianProduct(set)
	for pairs := range product.Iterator().C {
		first := pairs.(mapset.OrderedPair).First.(int64)
		second := pairs.(mapset.OrderedPair).Second.(int64)
		remainder := year - first - second

		if first != second && first != remainder && second != remainder && set.Contains(remainder) {
			return first * second * remainder
		}
	}
	return -1
}

func main() {
	file, _ := os.Open("input/day01.txt")
	defer file.Close()

	set := mapset.NewSet()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		newVal, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		set.Add(newVal)
	}
	println(reportRepairPart1(2020, set))
	println(reportRepairPart2(2020, set))
}
