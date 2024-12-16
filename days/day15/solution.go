package main

import (
	"fmt"
	"slices"
	"strings"
)

type coord struct {
	i, j int
}

var dirs = map[byte]coord{
	'>': {0, 1},
	'<': {0, -1},
	'^': {-1, 0},
	'v': {1, 0},
}

func part1(input string) any {
	gridS := strings.Split(input, "\n")
	ind := slices.Index(gridS, "")
	moves := strings.Join(gridS[ind+1:], "")
	gridS = gridS[:ind]
	grid := make([][]byte, len(gridS))
	for i := range grid {
		grid[i] = []byte(gridS[i])
	}

L:
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '@' {
				process(grid, coord{i, j}, moves)
				break L
			}
		}
	}

	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 'O' {
				res += i*100 + j
			}
		}
	}

	return res
}

func printGrid(grid [][]byte) {
	for _, row := range grid {
		fmt.Println(string(row))
	}
}

func process(grid [][]byte, start coord, moves string) {
	cur := start

	for i := 0; i < len(moves); i++ {
		// printGrid(grid)

		dir := dirs[moves[i]]
		next := coord{cur.i + dir.i, cur.j + dir.j}
		if grid[next.i][next.j] == '#' {
			continue
		}
		if grid[next.i][next.j] == '.' {
			grid[cur.i][cur.j], grid[next.i][next.j] = grid[next.i][next.j], grid[cur.i][cur.j]
			cur = next
			continue
		}

		tNext := next
		for {
			tNext = coord{tNext.i + dir.i, tNext.j + dir.j}
			if grid[tNext.i][tNext.j] == '.' {
				break
			}
			if grid[tNext.i][tNext.j] == '#' {
				break
			}
		}
		if grid[tNext.i][tNext.j] == '#' {
			continue
		}
		grid[cur.i][cur.j], grid[next.i][next.j] = grid[next.i][next.j], grid[cur.i][cur.j]
		grid[cur.i][cur.j], grid[tNext.i][tNext.j] = grid[tNext.i][tNext.j], grid[cur.i][cur.j]
		cur = next
	}
	// printGrid(grid)
}

func part2(input string) any {
	for _, line := range strings.Split(input, "\n") {
		_ = line
	}

	return nil
}
