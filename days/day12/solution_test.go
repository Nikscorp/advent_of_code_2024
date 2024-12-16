package main

import "testing"

// FIXME
const exampleInput1 string = `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`

const exampleInput1_2 string = `AAAA
BBCD
BBCC
EEEC`

const exampleInput1_3 string = `OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`

func TestPart1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  any
	}{
		{
			name:  "example",
			input: exampleInput1,
			// FIXME
			want: 1930,
		},
		{
			name:  "example2",
			input: exampleInput1_2,
			// FIXME
			want: 140,
		},
		{
			name:  "example2",
			input: exampleInput1_3,
			// FIXME
			want: 772,
		},
		// FIXME
		{
			name:  "real",
			input: input,
			want:  1352976,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

// FIXME
const exampleInput2 string = `AAAA
BBCD
BBCC
EEEC`

const exampleInput2_2 string = `AAAA
BBCD
BBCC
EEEC`

const exampleInput2_3 string = `EEEEE
EXXXX
EEEEE
EXXXX
EEEEE`

const exampleInput2_4 string = `AAAAAA
AAABBA
AAABBA
ABBAAA
ABBAAA
AAAAAA`

const exampleInput2_5 string = `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`

func TestPart2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  any
	}{
		{
			name:  "example",
			input: exampleInput2,
			// FIXME
			want: 80,
		},
		{
			name:  "example2",
			input: exampleInput2_2,
			// FIXME
			want: 436,
		},
		{
			name:  "example3",
			input: exampleInput2_3,
			// FIXME
			want: 236,
		},
		{
			name:  "example4",
			input: exampleInput2_4,
			// FIXME
			want: 368,
		},
		{
			name:  "example5",
			input: exampleInput2_5,
			// FIXME
			want: 1206,
		},
		// FIXME
		// {
		// 	name:  "real",
		// 	input: input,
		// 	want:  42,
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
