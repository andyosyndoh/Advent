package main

import (
	"fmt"
	"os"
	"strings"
)

var visited = make(map[string]bool)

var per, area = 0, 0

func main() {
	content, _ := os.ReadFile("instructions.txt")

	splited := strings.Split(string(content), "\n")
	tot := 0

	for i := range splited {
		for j := range splited[i] {
			if !visited[fmt.Sprintf("%v %v", i, j)] {
				per, area = 0, 0
				getRecursive(splited, i, j)
				tot += per * area
				// fmt.Println(string(splited[i][j]))
				// fmt.Println(per, area)
				// fmt.Println()
			}
		}
	}
	fmt.Println(tot)
}

var sides = [][]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func getRecursive(grid []string, x, y int) {
	visited[fmt.Sprintf("%v %v", x, y)] = true
	area++
	count := 0

	for _, v := range sides {
		nx, ny := x+v[0], y+v[1]
		if nx >= 0 && nx < len(grid) && ny >= 0 && ny < len(grid[x]) && grid[nx][ny] == grid[x][y] {
			count++
			if  !visited[fmt.Sprintf("%v %v", nx, ny)] {
				getRecursive(grid, nx, ny)
			}
		}
	}
	per += 4 - count
}
