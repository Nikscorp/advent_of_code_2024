package main

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	re  = regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	re2 = regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)
)

func part1(input string) any {
	res := 0
	for _, line := range strings.Split(input, "\n") {
		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			a, _ := strconv.Atoi(match[1])
			b, _ := strconv.Atoi(match[2])
			res += a * b
		}
	}

	return res
}

func part2(input string) any {
	res := 0
	enabled := true
	for _, line := range strings.Split(input, "\n") {
		matches := re2.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			if match[0] == "do()" {
				enabled = true
				continue
			}
			if match[0] == "don't()" {
				enabled = false
				continue
			}
			if !enabled {
				continue
			}
			a, _ := strconv.Atoi(match[1])
			b, _ := strconv.Atoi(match[2])
			res += a * b
		}
	}

	return res
}
