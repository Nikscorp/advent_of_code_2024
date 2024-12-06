package main

import (
	"strings"
)

type coord struct {
	i, j int
}

var dirs = []coord{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func part1(input string) any {
	grid := strings.Split(input, "\n")
	var start coord
L:
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '^' {
				start = coord{i, j}
				break L
			}
		}
	}

	visited := make(map[coord]struct{})
	visited[start] = struct{}{}
	indDir := 0
	cur := start

	for {
		dir := dirs[indDir]
		next := coord{cur.i + dir.i, cur.j + dir.j}
		if next.i < 0 || next.i >= len(grid) || next.j < 0 || next.j >= len(grid[0]) {
			break
		}
		if grid[next.i][next.j] == '#' {
			indDir = (indDir + 1) % len(dirs)
			continue
		}
		visited[next] = struct{}{}
		cur = next
	}

	return len(visited)
}

type coordWithDirInd struct {
	c      coord
	dirInd int
}

func part2(input string) any {
	gridS := strings.Split(input, "\n")
	grid := make([][]byte, len(gridS))
	for i, s := range gridS {
		grid[i] = []byte(s)
	}
	var start coord
L:
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '^' {
				start = coord{i, j}
				break L
			}
		}
	}

	res := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] != '.' {
				continue
			}
			grid[i][j] = '#'
			if hasLoop(grid, start) {
				res++
			}
			grid[i][j] = '.'
		}
	}

	return res
}

func hasLoop(grid [][]byte, start coord) bool {
	visited := make(map[coordWithDirInd]struct{})
	visited[coordWithDirInd{start, 0}] = struct{}{}
	indDir := 0
	cur := start
	loopFound := false

	for {
		dir := dirs[indDir]
		next := coord{cur.i + dir.i, cur.j + dir.j}
		if next.i < 0 || next.i >= len(grid) || next.j < 0 || next.j >= len(grid[0]) {
			break
		}
		if grid[next.i][next.j] == '#' {
			indDir = (indDir + 1) % len(dirs)
			continue
		}
		ci := coordWithDirInd{next, indDir}
		if _, ok := visited[ci]; ok {
			loopFound = true
			break
		}
		visited[ci] = struct{}{}
		cur = next
	}

	return loopFound
}
