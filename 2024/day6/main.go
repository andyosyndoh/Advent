package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	Lines  [][]rune     // Store the result grid
	Count  int          // Count of unsafe moves
	Unsafe int      = 1 // Tracks unsafe moves
)

func main() {
	// Read input file
	input, _ := os.ReadFile("instructions.txt")
	lines := strings.Split(string(input), "\n")

	// Convert lines to a 2D rune slice
	grid := make([][]rune, len(lines))
	for i := range lines {
		grid[i] = []rune(lines[i])
	}

	clone := cloneGrid(grid)

	spots := [][]int{}
	// Initialize
	here := make(map[string]rune)
	if solve(grid, here) {
		// for _, line := range Lines {
		// 	fmt.Println(string(line))
		// }
		spots = findSpots(Lines)
	}
	// Find starting position
	start := []int{}
	for i, row := range grid {
		for j, ch := range row {
			if ch == '^' {
				start = []int{i, j}
			}
		}
	}

	// Process unsafe moves
	c := 0
	for _, spot := range spots {
		if start[0] == spot[0] && start[1] == spot[1] {
			continue
		}
		c++
		here = make(map[string]rune)
		modifiedGrid := blockCell(clone, spot[0], spot[1])
		// for _, line := range modifiedGrid {
		// 	fmt.Println(string(line))
		// }
		// fmt.Println("1")
		if !solve(modifiedGrid, here) {
			// for _, line := range Lines {
			// 	fmt.Println(string(line))
			// }
			fmt.Println("Unsafe Spot:", spot)
			fmt.Println("Count:", c)
			Count++
		}
	}
	fmt.Println("Total Unsafe Moves:", Count)
}

// Blocks a specific cell in the grid by converting it to '#'
func blockCell(grid [][]rune, i, j int) [][]rune {
	newGrid := cloneGrid(grid)
	if grid[i][j] != '^' {
		newGrid[i][j] = '#'
	}
	return newGrid
}

// Solve function to recursively navigate the grid
func solve(grid [][]rune, Here map[string]rune) bool {
	rows, cols := len(grid), len(grid[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			ch := grid[i][j]
			if ch == '^' || ch == '>' || ch == '<' || ch == 'v' {
				pos := fmt.Sprintf("%v,%v",i,j) // Use integer as key
				if c, ok := Here[pos]; ok && c == ch {
					return false
				}
				Here[pos] = ch

				// Process movement
				switch ch {
				case '^':
					if i == 0 {
						Lines = cloneGrid(grid)
						return true
					} else if grid[i-1][j] != '#' && i > 0{
						grid[i][j], grid[i-1][j] = 'X', '^'
						if solve(grid, Here) {
							return true
						} else {
							return false
						}
						// grid[i][j], grid[i-1][j] = '^', '.' // Backtrack
					} else if grid[i-1][j] == '#' {
						grid[i][j] =  '>'
						if solve(grid, Here) {
							return true
						} else {
							return false
						}
						// grid[i][j] = '.' // Backtrack
					}
				case '>':
					if j == cols-1 {
						Lines = cloneGrid(grid)
						return true
					} else if grid[i][j+1] != '#' && j < len(grid[i])-1 {
						grid[i][j], grid[i][j+1] = 'X', '>'
						if solve(grid, Here) {
							return true
						} else {
							return false
						}
						// grid[i][j], grid[i][j+1] = '>', '.' // Backtrack
					} else if grid[i][j+1] == '#' {
						grid[i][j] = 'v'
						if solve(grid, Here) {
							return true
						} else {
							return false
						}
						// grid[i][j] = '.' // Backtrack
					}
				case 'v':
					if i == rows-1 {
						Lines = cloneGrid(grid)
						return true
					} else if grid[i+1][j] != '#' && i < len(grid)-1{
						grid[i][j], grid[i+1][j] = 'X', 'v'
						if solve(grid, Here) {
							return true
						} else {
							return false
						}
						// grid[i][j], grid[i+1][j] = 'v', '.' // Backtrack
					} else if grid[i+1][j] == '#' {
						grid[i][j] = '<'
						if solve(grid, Here) {
							return true
						} else {
							return false
						}
						// grid[i][j] = '.' // Backtrack
					}
				case '<':
					if j == 0 {
						Lines = cloneGrid(grid)
						return true
					} else if grid[i][j-1] != '#' && j > 0 {
						grid[i][j], grid[i][j-1] = 'X', '<'
						if solve(grid, Here) {
							return true
						} else {
							return false
						}
						// grid[i][j], grid[i][j-1] = '<', '.' // Backtrack
					} else if grid[i][j-1] == '#' {
						grid[i][j] =  '^'
						if solve(grid, Here) {
							return true
						} else {
							return false
						}
						// grid[i][j] = '.' // Backtrack
					}
				}
			}
		}
	}
	return false
}

// Clones a 2D grid to avoid modifying the original
func cloneGrid(grid [][]rune) [][]rune {
	cloned := make([][]rune, len(grid))
	for i := range grid {
		cloned[i] = append([]rune{}, grid[i]...)
	}
	return cloned
}

// Finds positions of interest in the grid
func findSpots(grid [][]rune) [][]int {
	spots := [][]int{}
	for i, row := range grid {
		for j, ch := range row {
			if ch == 'X' || ch == '^' || ch == '>' || ch == '<' || ch == 'v' {
				spots = append(spots, []int{i, j})
			}
		}
	}
	fmt.Println(len(spots))
	return spots
}
