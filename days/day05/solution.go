package main

import (
	"strconv"
	"strings"
)

func part1(input string) any {
	lines := strings.Split(input, "\n")
	ind := 0
	ab := make(map[int][]int)
	res := 0

	for ; lines[ind] != ""; ind++ {
		line := lines[ind]
		as, bs, _ := strings.Cut(line, "|")
		a, _ := strconv.Atoi(as)
		b, _ := strconv.Atoi(bs)
		ab[a] = append(ab[a], b)
	}

	for _, line := range lines[ind+1:] {
		numsS := strings.Split(line, ",")
		nums := make([]int, 0, len(numsS))
		indeces := make(map[int]int, len(numsS))
		for i, numS := range numsS {
			num, _ := strconv.Atoi(numS)
			nums = append(nums, num)
			indeces[num] = i
		}

		if len(indeces) != len(nums) {
			panic("nums are not unique")
		}

		isValid := true
		for numIndex, num := range nums {
			numsAfter, ok := ab[num]
			if !ok {
				continue
			}

			for _, numAfter := range numsAfter {
				afterIndex, ok := indeces[numAfter]
				if !ok {
					continue
				}

				if afterIndex <= numIndex {
					isValid = false
					break
				}
			}

			if !isValid {
				break
			}
		}

		if isValid {
			res += nums[len(nums)/2]
		}
	}

	return res
}

func part2(input string) any {
	lines := strings.Split(input, "\n")
	ind := 0
	ab := make(map[int][]int)
	res := 0

	for ; lines[ind] != ""; ind++ {
		line := lines[ind]
		as, bs, _ := strings.Cut(line, "|")
		a, _ := strconv.Atoi(as)
		b, _ := strconv.Atoi(bs)
		ab[a] = append(ab[a], b)
	}

	for _, line := range lines[ind+1:] {
		numsS := strings.Split(line, ",")
		nums := make([]int, 0, len(numsS))
		indeces := make(map[int]int, len(numsS))
		for i, numS := range numsS {
			num, _ := strconv.Atoi(numS)
			nums = append(nums, num)
			indeces[num] = i
		}

		if len(indeces) != len(nums) {
			panic("nums are not unique")
		}

		wasInvalid := false

		for {
			isValid := true
			for i, num := range nums {
				indeces[num] = i
			}
			for numIndex, num := range nums {
				numsAfter, ok := ab[num]
				if !ok {
					continue
				}

				for _, numAfter := range numsAfter {
					afterIndex, ok := indeces[numAfter]
					if !ok {
						continue
					}

					if afterIndex <= numIndex {
						isValid = false
						wasInvalid = true
						nums[afterIndex], nums[numIndex] = nums[numIndex], nums[afterIndex]
						break
					}
				}

				if !isValid {
					break
				}
			}

			if isValid {
				break
			}
		}

		if wasInvalid {
			res += nums[len(nums)/2]
		}
	}

	return res
}
