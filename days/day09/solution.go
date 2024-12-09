package main

import (
	"slices"
	"strings"
)

func part1(input string) any {
	res := make([]int, 0)
	line := strings.TrimSpace(input)
	curInd := 0
	freePos := -1

	for i := 0; i < len(line); i += 2 {
		n := int(line[i] - '0')
		free := 0
		if i+1 < len(line) {
			free = int(line[i+1] - '0')
		}

		for range n {
			res = append(res, curInd)
		}
		for range free {
			if freePos == -1 {
				freePos = len(res)
			}
			res = append(res, -1)
		}
		curInd++
	}

	curPos := len(res) - 1
	for freePos < len(res) && freePos < curPos {
		res[freePos] = res[curPos]
		res[curPos] = -1
		for freePos < len(res) && res[freePos] != -1 {
			freePos++
		}
		for curPos > freePos && res[curPos] == -1 {
			curPos--
		}
	}

	checksum := 0
	for i, n := range res {
		if n == -1 {
			continue
		}
		checksum += i * n
	}

	return checksum
}

type block struct {
	fileID int
	cnt    int
}

func part2(input string) any {
	res := make([]block, 0)
	line := strings.TrimSpace(input)
	curInd := 0

	for i := 0; i < len(line); i += 2 {
		n := int(line[i] - '0')
		res = append(res, block{curInd, n})
		free := 0
		if i+1 < len(line) {
			free = int(line[i+1] - '0')
			res = append(res, block{-1, free})
		}

		curInd++
	}

	for fileID := curInd - 1; fileID >= 0; fileID-- {
		blockInd := slices.IndexFunc(res, func(x block) bool {
			return x.fileID == fileID
		})
		mblock := res[blockInd]

		freeInd := slices.IndexFunc(res, func(x block) bool {
			return x.fileID == -1 && x.cnt >= mblock.cnt
		})
		if freeInd == -1 {
			continue
		}
		if freeInd >= blockInd {
			continue
		}
		res[blockInd].fileID = -1
		res[freeInd].cnt -= mblock.cnt
		if res[freeInd].cnt == 0 {
			res[freeInd] = mblock
			continue
		}
		res = slices.Insert(res, freeInd, mblock)
	}

	checksum := 0
	t := 0
	for _, mblock := range res {
		for range mblock.cnt {
			if mblock.fileID != -1 {
				checksum += t * mblock.fileID
			}
			t++
		}
	}

	return checksum
}
