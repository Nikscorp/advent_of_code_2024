package main

import (
	"math"
	"math/big"
	"regexp"
	"strconv"
	"strings"
)

var (
	buttonRe = regexp.MustCompile(`Button (A|B): X\+(\d+), Y\+(\d+)`)
	prizeRe  = regexp.MustCompile(`Prize: X=(\d+), Y=(\d+)`)
)

const (
	aCost = 3
	bCost = 1
)

type coord struct {
	x, y int
}

func part1(input string) any {
	lines := strings.Split(input, "\n")
	res := 0

	for i := 0; i < len(lines); i += 4 {
		matchesA := buttonRe.FindStringSubmatch(lines[i])
		adx, _ := strconv.Atoi(matchesA[2])
		ady, _ := strconv.Atoi(matchesA[3])
		ac := coord{adx, ady}

		matchesB := buttonRe.FindStringSubmatch(lines[i+1])
		bdx, _ := strconv.Atoi(matchesB[2])
		bdy, _ := strconv.Atoi(matchesB[3])
		bc := coord{bdx, bdy}

		matchesP := prizeRe.FindStringSubmatch(lines[i+2])
		px, _ := strconv.Atoi(matchesP[1])
		py, _ := strconv.Atoi(matchesP[2])
		pc := coord{px, py}

		cur := dp(coord{0, 0}, ac, bc, pc, make(map[coord]int))
		// cur := gauss(ac, bc, pc)
		if cur != -1 {
			res += cur
		}
	}

	return res
}

func dp(sc, ac, bc, pc coord, memo map[coord]int) int {
	if sc == pc {
		return 0
	}

	if sc.x > pc.x || sc.y > pc.y {
		return -1
	}

	if v, ok := memo[sc]; ok {
		return v
	}

	var byA, byB *int
	tryA := dp(coord{sc.x + ac.x, sc.y + ac.y}, ac, bc, pc, memo)
	if tryA != -1 {
		tryA += aCost
		byA = &tryA
	}

	tryB := dp(coord{sc.x + bc.x, sc.y + bc.y}, ac, bc, pc, memo)
	if tryB != -1 {
		tryB += bCost
		byB = &tryB
	}

	if byA == nil && byB == nil {
		memo[sc] = -1
		return -1
	}

	res := math.MaxInt
	if byA != nil {
		res = min(res, *byA)
	}
	if byB != nil {
		res = min(res, *byB)
	}

	memo[sc] = res
	return res
}

func gauss(ac, bc, pc coord) int {
	res := 0

	m := make([][]*big.Rat, 2)
	m[0] = []*big.Rat{big.NewRat(int64(ac.x), 1), big.NewRat(int64(bc.x), 1), big.NewRat(int64(pc.x), 1)}
	m[1] = []*big.Rat{big.NewRat(int64(ac.y), 1), big.NewRat(int64(bc.y), 1), big.NewRat(int64(pc.y), 1)}

	div := big.NewRat(1, 1)
	div.Set(m[1][0])

	m[1][0].Quo(m[1][0], div)
	m[1][1].Quo(m[1][1], div)
	m[1][2].Quo(m[1][2], div)

	mul := big.NewRat(1, 1)
	mul.Set(m[0][0])

	cpy := big.NewRat(1, 1)
	m[0][0].Sub(m[0][0], cpy.Mul(m[1][0], mul))
	m[0][1].Sub(m[0][1], cpy.Mul(m[1][1], mul))
	m[0][2].Sub(m[0][2], cpy.Mul(m[1][2], mul))

	pressB := cpy.Quo(m[0][2], m[0][1])
	if !pressB.IsInt() {
		return -1
	}

	res += int(pressB.Num().Int64()) * bCost

	cpy.Mul(pressB, m[1][1])
	cpy.Sub(m[1][2], cpy)
	pressA := cpy.Quo(cpy, m[1][0])
	if !pressB.IsInt() {
		return -1
	}

	res += int(pressA.Num().Int64()) * aCost

	return res
}

func part2(input string) any {
	lines := strings.Split(input, "\n")
	res := 0

	for i := 0; i < len(lines); i += 4 {
		matchesA := buttonRe.FindStringSubmatch(lines[i])
		adx, _ := strconv.Atoi(matchesA[2])
		ady, _ := strconv.Atoi(matchesA[3])
		ac := coord{adx, ady}

		matchesB := buttonRe.FindStringSubmatch(lines[i+1])
		bdx, _ := strconv.Atoi(matchesB[2])
		bdy, _ := strconv.Atoi(matchesB[3])
		bc := coord{bdx, bdy}

		matchesP := prizeRe.FindStringSubmatch(lines[i+2])
		px, _ := strconv.Atoi(matchesP[1])
		py, _ := strconv.Atoi(matchesP[2])
		pc := coord{px + 10000000000000, py + 10000000000000}

		cur := gauss(ac, bc, pc)
		if cur != -1 {
			res += cur
		}
	}

	return res
}
