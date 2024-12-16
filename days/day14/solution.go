package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type coord struct {
	i, j int
}

var re = regexp.MustCompile(`p=(-*\d+),(-*\d+) v=(-*\d+),(-*\d+)`)

const (
	// maxI = 7
	// maxJ = 11
	maxI = 103
	maxJ = 101

	nSecs = 100
)

func part1(input string) any {
	res := make([]coord, 0)
	for _, line := range strings.Split(input, "\n") {
		matches := re.FindStringSubmatch(line)
		pi, _ := strconv.Atoi(matches[1])
		pj, _ := strconv.Atoi(matches[2])
		vi, _ := strconv.Atoi(matches[3])
		vj, _ := strconv.Atoi(matches[4])
		p := coord{pj, pi}
		v := coord{vj, vi}
		cur := coord{((p.i+nSecs*v.i)%maxI + maxI) % maxI, (((p.j + nSecs*v.j) % maxJ) + maxJ) % maxJ}
		res = append(res, cur)
	}

	var one, two, three, four int
	for _, c := range res {
		if c.i < maxI/2 {
			if c.j < maxJ/2 {
				one++
			} else if c.j > maxJ/2 {
				two++
			}
		} else if c.i > maxI/2 {
			if c.j < maxJ/2 {
				three++
			} else if c.j > maxJ/2 {
				four++
			}
		}
	}

	return one * two * three * four
}

func part2(input string) any {
	res := make([][]coord, 0)
	for j := 0; j < 100000; j++ {
		res = append(res, make([]coord, 0))
	}

	for _, line := range strings.Split(input, "\n") {
		matches := re.FindStringSubmatch(line)
		pi, _ := strconv.Atoi(matches[1])
		pj, _ := strconv.Atoi(matches[2])
		vi, _ := strconv.Atoi(matches[3])
		vj, _ := strconv.Atoi(matches[4])
		p := coord{pj, pi}
		v := coord{vj, vi}

		for j := 0; j < 100000; j++ {
			cur := coord{((p.i+j*v.i)%maxI + maxI) % maxI, (((p.j + j*v.j) % maxJ) + maxJ) % maxJ}
			res[j] = append(res[j], cur)
		}
	}

	for j := 0; j < 100000; j++ {
		print(j, res[j])
	}

	return 0
}

func print(j int, res []coord) {
	m := make(map[coord]bool)
	for _, c := range res {
		m[c] = true
	}
	if len(m) != len(res) {
		return
	}
	fmt.Println(j)
	for i := 0; i < maxI; i++ {
		for j := 0; j < maxJ; j++ {
			if m[coord{i, j}] {
				fmt.Printf("X")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
}
