package main

import (
	"slices"
	"strconv"
	"strings"
)

func part1(input string) any {
	arr1 := make([]int, 0)
	arr2 := make([]int, 0)
	for _, line := range strings.Split(input, "\n") {
		s1, s2, _ := strings.Cut(line, "   ")
		n1, _ := strconv.Atoi(s1)
		n2, _ := strconv.Atoi(s2)
		arr1 = append(arr1, n1)
		arr2 = append(arr2, n2)
	}

	slices.Sort(arr1)
	slices.Sort(arr2)

	res := 0
	for i := range arr1 {
		res += abs(arr1[i] - arr2[i])
	}

	return res
}

func part2(input string) any {
	arr1 := make([]int, 0)
	m2 := make(map[int]int)
	for _, line := range strings.Split(input, "\n") {
		s1, s2, _ := strings.Cut(line, "   ")
		n1, _ := strconv.Atoi(s1)
		n2, _ := strconv.Atoi(s2)
		arr1 = append(arr1, n1)
		m2[n2]++
	}

	res := 0
	for _, n := range arr1 {
		res += n * m2[n]
	}

	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
