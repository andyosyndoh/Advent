package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("instructions.txt")

	lines := strings.Split(string(input), "\n")

	grid := make([][]rune, len(lines))
	

	for i, line := range lines {
		grid[i] = []rune(line)
	}
	n := 0
	fin := 0
	x, y := len(grid[0]), len(grid)
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if check(grid[i][j]) {
				grid, n = addnodes(i, j, grid)
				fin += n
			}
		}
	}

	count := 0
	for _, l := range grid {
		for _, ch := range l {
			if ch != '.' {
				count++
			}
		}
	}

	for _, l := range grid {
		fmt.Println(string(l))
	}
	fmt.Println(count, fin)
}

func check(ch rune) bool {
	if ch == '.' || ch == '#' {
		return false
	}
	return true
}

func addnodes(i, j int, grid [][]rune) ([][]rune, int) {
	count := 0
	ch := grid[i][j]
	row, col := len(grid), len(grid[0])
	for di := 0; di < row; di++ {
		for dj := 0; dj < col; dj++ {
			if grid[di][dj] == ch {
				disI1, disJ1 := (di - i), (dj - j)
				num := 1
				for con(i, j, di, dj) {
					num += 1
					disI, disJ := disI1*num, disJ1*num
					// fmt.Println(i+disI, j+disJ)
					if i+disI < 0 || j+disJ < 0 || i+disI > row-1 || j+disJ > col-1 {
						break
					} else if grid[i+disI][j+disJ] == '.' {
						grid = place(i+disI, j+disJ, grid)
						count++
					} else if grid[i+disI][j+disJ] != '.' {
						if (grid[i+disI][j+disJ]) != '#' {
							fmt.Println(string((grid[i+disI][j+disJ])))
							count++
						}
					}
				}
			}
		}
	}
	return grid, count
}

func con(i, j, di, dj int) bool {
	if i == di && j == dj {
		return false
	}
	return true
}

func place(i, j int, grid [][]rune) [][]rune {
	var new [][]rune
	for x := range grid {
		var n []rune
		for y := range grid[x] {
			if x == i && y == j {
				n = append(n, '#')
			} else {
				n = append(n, grid[x][y])
			}
		}
		new = append(new, n)
	}
	return new
}
