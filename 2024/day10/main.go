package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var directions = [][2]int{
	{-1, 0}, // up
	{1, 0},  // down
	{0, -1}, // left
	{0, 1},  // right
}

func main() {
	input, _ := os.ReadFile("instructions.txt")

	grid := strings.Split(string(input), "\n")

	var paths [][]string

	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}

	for i, line := range grid {
		for j, ch := range line {
			if ch == '0' {
				// fmt.Println(i, j)

				findpaths(grid, i, j, visited, []string{}, &paths)
			}
		}
	}

	// for _, f := range paths {
	// 	fmt.Println(f)
	// 	// fmt.Println("hello")
	// }

	fmt.Println(len(paths))
}

func findpaths(grid []string, i, j int, visited [][]bool, path []string, paths *[][]string) {
	visited[i][j] = true
	path = append(path, fmt.Sprintf("%d,%d", i, j))

	if grid[i][j] == '9' {
		f := make([]string, len(path))
		copy(f, path)
		// if comp(*paths, f) {
		*paths = append(*paths, f)
		// }
		visited[i][j] = false
		return
	}

	for _, dir := range directions {
		ni, nj := i+dir[0], j+dir[1]
		// fmt.Println(ni,nj)
		if ni >= 0 && nj >= 0 && ni < len(grid) && nj < len(grid[ni]) && !visited[ni][nj] && grid[ni][nj] != '.' {
			curr, _ := strconv.Atoi(string(grid[i][j]))
			next, _ := strconv.Atoi(string(grid[ni][nj]))
			if curr+1 == next {
				// fmt.Print(path)
				findpaths(grid, ni, nj, visited, path, paths)
			}
		}
	}

	visited[i][j] = false
}

func comp(paths [][]string, path []string) bool {
	count := 0
	for _, p := range paths {
		for i := range p {
			if p[i] != path[i] {
				count++
			}
		}
		if count == 9 {
			return false
		}
		count = 0
	}
	return true
}
