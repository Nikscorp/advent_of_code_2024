package main

import "strings"

type coord struct {
	i, j int
}

var dirs = []coord{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func part1(input string) any {
	grid := strings.Split(input, "\n")
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '0' {
				res += bfs(grid, coord{i, j}, true)
			}
		}
	}

	return res
}

func bfs(grid []string, start coord, considerSeen bool) int {
	seen := make(map[coord]bool)
	res := 0
	seen[start] = true
	queue := []coord{start}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		for _, dir := range dirs {
			next := coord{cur.i + dir.i, cur.j + dir.j}
			if next.i < 0 || next.i >= len(grid) || next.j < 0 || next.j >= len(grid[0]) {
				continue
			}

			if grid[next.i][next.j]-grid[cur.i][cur.j] != 1 {
				continue
			}

			if seen[next] {
				continue
			}
			if considerSeen {
				seen[next] = true
			}
			if grid[next.i][next.j] == '9' {
				res++
				continue
			}
			queue = append(queue, next)
		}
	}

	return res
}

func part2(input string) any {
	grid := strings.Split(input, "\n")
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '0' {
				res += bfs(grid, coord{i, j}, false)
			}
		}
	}

	return res
}
