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

	visited := make(map[string]bool)

	var paths [][]string

	for i, line := range grid {
		for j, ch := range line {
			if ch == '0' {
				fmt.Println(i,j)
				findpaths(grid, i, j, visited, []string{}, &paths)
			}
		}
	}

	for _, f := range paths {
		// for _, g := range f {
		// 	s := strings.Split(g, ",")
		// 	f, _ := strconv.Atoi(s[0])
		// 	n, _ := strconv.Atoi(s[1])
		// 	fmt.Print(grid[f][n])
		// 	fmt.Print(" ")
		// }
		fmt.Println(f)
	}

	fmt.Println(len(paths))
}

func findpaths(grid []string, i, j int, visited map[string]bool, path []string, paths *[][]string) {
	visited[fmt.Sprintf("%d,%d", i, j)] = true
	path = append(path, fmt.Sprintf("%d,%d", i, j))

	if grid[i][j] == '9' {
		if conf(*paths, path) {
			*paths = append(*paths, path)
		}
		visited[fmt.Sprintf("%d,%d", i, j)] = false
		return
	}

	for _, dir := range directions {
		ni, nj := i+dir[0], j+dir[1]
		if ni >= 0 && nj >= 0 && ni < len(grid) && nj < len(grid[0]) && !visited[fmt.Sprintf("%d,%d", i, j)] && grid[ni][nj] != '.' {
			curr, _ := strconv.Atoi(string(grid[i][j]))
			next, _ := strconv.Atoi(string(grid[ni][nj]))
			if curr+1 == next {
				findpaths(grid, ni, nj, visited, path, paths)
			}
		}
	}

	visited[fmt.Sprintf("%d,%d", i, j)] = false
	// path = path[:len(path)-1]
}

func comp(back, path []string) bool {
	for i:= 0; i < 9; i++ {
		if back[i] != path[i] {
			return true
		}
	}
	return false
}

func conf(paths [][]string, path []string) bool {
	for _, back := range paths {
		if !comp(back, path) {
			return false
		}
	}
	return true
}