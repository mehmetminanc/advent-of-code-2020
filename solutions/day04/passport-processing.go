package main

import (
	"bufio"
	"bytes"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input/day04.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(ScanDoubleLines)
	countP1 := 0
	countP2 := 0
	for scanner.Scan() {
		text := scanner.Text()
		passportMap := make(map[string]string)
		passport := strings.ReplaceAll(text, "\n", " ")
		parts := strings.Split(passport, " ")
		for _, part := range parts {
			kv := strings.Split(part, ":")
			if len(kv) > 1 {
				passportMap[kv[0]] = kv[1]
			}
		}
		if isValidForPart1(passportMap) {
			countP1 += 1
			if isValidForPart2(passportMap) {
				countP2 += 1
			}
		}
	}
	println(countP1)
	println(countP2)
}

func isValidForPart1(passportMap map[string]string) bool {
	for _, key := range []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"} {
		_, ok := passportMap[key]
		if !ok {
			return false
		}
	}
	return true
}

func isValidForPart2(passportMap map[string]string) bool {
	byr, _ := strconv.ParseInt(passportMap["byr"], 10, 64)
	iyr, _ := strconv.ParseInt(passportMap["iyr"], 10, 64)
	eyr, _ := strconv.ParseInt(passportMap["eyr"], 10, 64)
	matchCm, _ := regexp.MatchString("^((15\\d)|(16\\d)|(17\\d)|(18\\d)|(19[0123]))cm$", passportMap["hgt"])
	matchIn, _ := regexp.MatchString("^(59|6\\d|7[0123456])in$", passportMap["hgt"])
	matchHcl, _ := regexp.MatchString("^#[0-9a-f]{6}$", passportMap["hcl"])
	matchEcl, _ := regexp.MatchString("^(amb|blu|brn|gry|grn|hzl|oth)$", passportMap["ecl"])
	matchPid, _ := regexp.MatchString("^\\d{9}$", passportMap["pid"])

	if byr >= 1920 && byr <= 2002 {
		if iyr >= 2010 && iyr <= 2020 {
			if eyr >= 2020 && eyr <= 2030 {
				if matchCm || matchIn {
					if matchHcl {
						if matchEcl {
							if matchPid {
								return true
							}
						}
					}
				}
			}
		}
	}
	return false
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
