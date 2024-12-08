package main

import "strings"

type coord struct {
	i, j int
}

func part1(input string) any {
	grid := strings.Split(input, "\n")
	freqToAntennas := make(map[byte][]coord)

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '.' {
				continue
			}
			freqToAntennas[grid[i][j]] = append(freqToAntennas[grid[i][j]], coord{i, j})
		}
	}

	seen := make(map[coord]struct{})
	for _, antennas := range freqToAntennas {
		for i := 0; i < len(antennas); i++ {
			for j := i + 1; j < len(antennas); j++ {
				a1 := antennas[i]
				a2 := antennas[j]
				delta := coord{a1.i - a2.i, a1.j - a2.j}
				c1 := coord{a1.i + delta.i, a1.j + delta.j}
				c2 := coord{a2.i - delta.i, a2.j - delta.j}
				if c1.i >= 0 && c1.i < len(grid) && c1.j >= 0 && c1.j < len(grid[0]) {
					seen[c1] = struct{}{}
				}
				if c2.i >= 0 && c2.i < len(grid) && c2.j >= 0 && c2.j < len(grid[0]) {
					seen[c2] = struct{}{}
				}
			}
		}
	}

	return len(seen)
}

func part2(input string) any {
	grid := strings.Split(input, "\n")
	freqToAntennas := make(map[byte][]coord)

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '.' {
				continue
			}
			freqToAntennas[grid[i][j]] = append(freqToAntennas[grid[i][j]], coord{i, j})
		}
	}

	seen := make(map[coord]struct{})
	for _, antennas := range freqToAntennas {
		for i := 0; i < len(antennas); i++ {
			for j := i + 1; j < len(antennas); j++ {
				a1 := antennas[i]
				a2 := antennas[j]
				seen[a1] = struct{}{}
				seen[a2] = struct{}{}
				delta := coord{a1.i - a2.i, a1.j - a2.j}
				for {
					c1 := coord{a1.i + delta.i, a1.j + delta.j}
					if c1.i >= 0 && c1.i < len(grid) && c1.j >= 0 && c1.j < len(grid[0]) {
						seen[c1] = struct{}{}
						a1 = c1
					} else {
						break
					}
				}
				for {
					c2 := coord{a2.i - delta.i, a2.j - delta.j}
					if c2.i >= 0 && c2.i < len(grid) && c2.j >= 0 && c2.j < len(grid[0]) {
						seen[c2] = struct{}{}
						a2 = c2
					} else {
						break
					}
				}
			}
		}
	}

	return len(seen)
}
