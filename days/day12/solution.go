package main

import "strings"

type coord struct {
	i, j int
}

var dirs = []coord{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

func part1(input string) any {
	grid := strings.Split(input, "\n")
	seen := make(map[coord]bool)
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			c := coord{i, j}
			if seen[c] {
				continue
			}
			res += bfs(grid, seen, c)
		}
	}

	return res
}

func bfs(grid []string, seen map[coord]bool, start coord) int {
	area := 1
	perimeter := 0

	seen[start] = true
	queue := []coord{start}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		curPerimiter := 4

		for _, dir := range dirs {
			next := coord{cur.i + dir.i, cur.j + dir.j}
			if next.i < 0 || next.i >= len(grid) || next.j < 0 || next.j >= len(grid[0]) {
				continue
			}
			if grid[next.i][next.j] != grid[start.i][start.j] {
				continue
			}

			curPerimiter--
			if seen[next] {
				continue
			}
			seen[next] = true
			queue = append(queue, next)
			area += 1
		}

		perimeter += curPerimiter
	}

	return perimeter * area
}

func part2(input string) any {
	grid := strings.Split(input, "\n")
	seen := make(map[coord]bool)
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			c := coord{i, j}
			if seen[c] {
				continue
			}
			res += bfs2(grid, seen, c)
		}
	}

	return res
}

type coordWithDir struct {
	c   coord
	dir coord
}

func bfs2(grid []string, seen map[coord]bool, start coord) int {
	area := 1
	numberOfSides := 0

	seen[start] = true
	queue := []coordWithDir{{start, coord{-1, -1}}}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		angles := 0
		for _, dir := range dirs {
			next := coord{cur.c.i + dir.i, cur.c.j + dir.j}
			if next.i < 0 || next.i >= len(grid) || next.j < 0 || next.j >= len(grid[0]) {
				continue
			}
			if grid[next.i][next.j] != grid[start.i][start.j] {
				continue
			}

			if seen[next] {
				continue
			}
			seen[next] = true
			queue = append(queue, coordWithDir{next, dir})
			area += 1
			if cur.dir != dir {
				angles++
			}
		}

		numberOfSides += angles
	}

	return numberOfSides * area
}
