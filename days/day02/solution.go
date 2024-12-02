package main

import (
	"strconv"
	"strings"
)

func part1(input string) any {
	res := 0
	for _, line := range strings.Split(input, "\n") {
		arrS := strings.Fields(line)
		arr := make([]int, len(arrS))
		for i := range arrS {
			arr[i], _ = strconv.Atoi(arrS[i])
		}

		if isSafe(arr) {
			res++
		}
	}

	return res
}

func isSafe(arr []int) bool {
	allDec := true
	allInc := true
	okDiff := true
	for i := range len(arr) - 1 {
		if arr[i] >= arr[i+1] {
			allInc = false
		}
		if arr[i] <= arr[i+1] {
			allDec = false
		}
		diff := abs(arr[i] - arr[i+1])
		if diff < 1 || diff > 3 {
			okDiff = false
		}
		if !(allDec || allInc || okDiff) {
			break
		}
	}

	return allDec && okDiff || allInc && okDiff
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func part2(input string) any {
	res := 0
	for _, line := range strings.Split(input, "\n") {
		arrS := strings.Fields(line)
		arr := make([]int, len(arrS))
		for i := range arrS {
			arr[i], _ = strconv.Atoi(arrS[i])
		}

		if isSafe(arr) {
			res++
			continue
		}

		for i := range arr {
			tmpArr := make([]int, 0, len(arr)-1)
			tmpArr = append(tmpArr, arr[:i]...)
			tmpArr = append(tmpArr, arr[i+1:]...)
			if isSafe(tmpArr) {
				res++
				break
			}
		}
	}

	return res
}
