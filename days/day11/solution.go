package main

import (
	"strconv"
	"strings"
)

func part1(input string) any {
	stonesStr := strings.Fields(strings.TrimSpace(input))
	stones := make(map[int]int, len(stonesStr))
	for _, stoneStr := range stonesStr {
		stone, _ := strconv.Atoi(stoneStr)
		stones[stone]++
	}

	for range 25 {
		newStones := make(map[int]int, len(stones))
		for id, cnt := range stones {
			if id == 0 {
				newStones[1] += cnt
				continue
			}

			if v := digitsCnt(id); v%2 == 0 {
				first, second := digitsSplit(id, v)
				newStones[first] += cnt
				newStones[second] += cnt
				continue
			}

			newStones[id*2024] += cnt
		}
		stones = newStones
	}

	res := 0
	for _, v := range stones {
		res += v
	}

	return res
}

func digitsCnt(n int) int {
	res := 0
	for n > 0 {
		res++
		n /= 10
	}
	return res
}

func digitsSplit(n int, k int) (int, int) {
	second := 0
	first := n
	t := 1

	for range k / 2 {
		second = second + first%10*t
		first /= 10
		t *= 10
	}

	return first, second
}

func part2(input string) any {
	stonesStr := strings.Fields(strings.TrimSpace(input))
	stones := make(map[int]int, len(stonesStr))
	for _, stoneStr := range stonesStr {
		stone, _ := strconv.Atoi(stoneStr)
		stones[stone]++
	}

	for range 75 {
		newStones := make(map[int]int, len(stones))
		for id, cnt := range stones {
			if id == 0 {
				newStones[1] += cnt
				continue
			}

			if v := digitsCnt(id); v%2 == 0 {
				first, second := digitsSplit(id, v)
				newStones[first] += cnt
				newStones[second] += cnt
				continue
			}

			newStones[id*2024] += cnt
		}
		stones = newStones
	}

	res := 0
	for _, v := range stones {
		res += v
	}

	return res
}
