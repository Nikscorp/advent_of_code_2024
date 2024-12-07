package main

import (
	"slices"
	"strconv"
	"strings"
)

func part1(input string) any {
	res := 0
	for _, line := range strings.Split(input, "\n") {
		colonInd := strings.Index(line, ":")
		wantS := line[:colonInd]
		operandsS := strings.Fields(line[colonInd+2:])
		operands := make([]int, 0, len(operandsS))
		for _, operandS := range operandsS {
			operand, _ := strconv.Atoi(operandS)
			operands = append(operands, operand)
		}
		want, _ := strconv.Atoi(wantS)

		curs := []int{operands[0]}
		for i := 1; i < len(operands); i++ {
			nexts := make([]int, 0, len(curs)*2)
			for _, cur := range curs {
				nexts = append(nexts, cur+operands[i], cur*operands[i])
			}
			curs = nexts
		}

		if slices.Contains(curs, want) {
			res += want
		}
	}

	return res
}

func part2(input string) any {
	res := 0
	for _, line := range strings.Split(input, "\n") {
		colonInd := strings.Index(line, ":")
		wantS := line[:colonInd]
		operandsS := strings.Fields(line[colonInd+2:])
		operands := make([]int, 0, len(operandsS))
		for _, operandS := range operandsS {
			operand, _ := strconv.Atoi(operandS)
			operands = append(operands, operand)
		}
		want, _ := strconv.Atoi(wantS)

		curs := []int{operands[0]}
		for i := 1; i < len(operands); i++ {
			nexts := make([]int, 0, len(curs)*3)
			for _, cur := range curs {
				tcur := cur
				toperand := operands[i]
				for toperand > 0 {
					tcur *= 10
					toperand /= 10
				}
				tcur += operands[i]
				nexts = append(nexts, cur+operands[i], cur*operands[i], tcur)
			}
			curs = nexts
		}

		if slices.Contains(curs, want) {
			res += want
		}
	}

	return res
}
