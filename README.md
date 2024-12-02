# Advent of code 2024

<div align="center">

![](https://img.shields.io/badge/day%20📅-2-blue) ![](https://img.shields.io/badge/stars%20⭐-4-yellow) ![](https://img.shields.io/badge/days%20completed-2-red)

<img src="./static/logo.svg" width="220" />

</div>

Here you can find my golang solutions for [Advent Of Code 2024](https://adventofcode.com).

<!--- advent_readme_stars table --->
## 2024 Results

| Day | Part 1 | Part 2 |
| :---: | :---: | :---: |
| [Day 1](https://adventofcode.com/2024/day/1) | 🌟 | 🌟 |
| [Day 2](https://adventofcode.com/2024/day/2) | 🌟 | 🌟 |
<!--- advent_readme_stars table --->

## Codegen

`days/day*` codegen is a slightly changed version of [alexchao26](https://github.com/alexchao26/advent-of-code-go) aoc repo.

### Workflow

To generate skeleton for another day run:

```console
$ make gen DAY=3
```

Then manually:
1. Fill tests with example input and want values
2. Fill input.txt with real data
3. Code problem solution

Another commands:

```console
$ make run DAY=03 PART=2
$ make test DAY=03
```
