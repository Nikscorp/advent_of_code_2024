package main

import "strings"

type coord struct {
	i, j int
}

var dirs = []coord{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
	{1, 1},
	{-1, -1},
	{1, -1},
	{-1, 1},
}

func part1(input string) any {
	res := 0
	grid := strings.Split(input, "\n")
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			res += count(grid, coord{i, j}, "XMAS")
		}
	}

	return res
}

func count(grid []string, start coord, want string) int {
	res := 0
	if grid[start.i][start.j] != want[0] {
		return 0
	}
	wantBkp := want[1:]

	for _, dir := range dirs {
		cur := start
		want = wantBkp

		for {
			newC := coord{cur.i + dir.i, cur.j + dir.j}
			if newC.i < 0 || newC.i >= len(grid) || newC.j < 0 || newC.j >= len(grid[0]) {
				break
			}
			if grid[newC.i][newC.j] != want[0] {
				break
			}
			cur = newC
			want = want[1:]
			if want == "" {
				res++
				break
			}
		}
	}

	return res
}

func part2(input string) any {
	res := 0
	grid := strings.Split(input, "\n")

	for i := 0; i < len(grid)-2; i++ {
		for j := 0; j < len(grid)-2; j++ {
			if grid[i+1][j+1] != 'A' {
				continue
			}
			wantC := 'z'
			if grid[i][j] == 'S' {
				wantC = 'M'
			} else if grid[i][j] == 'M' {
				wantC = 'S'
			} else {
				continue
			}
			if grid[i+2][j+2] != byte(wantC) {
				continue
			}

			if grid[i+2][j] == 'S' {
				wantC = 'M'
			} else if grid[i+2][j] == 'M' {
				wantC = 'S'
			} else {
				continue
			}

			if grid[i][j+2] != byte(wantC) {
				continue
			}
			res++
		}
	}

	return res
}
