package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	input, _ := os.ReadFile("instructions.txt")

	new := []string{}

	count := 0
	for i, ch := range string(input) {
		if i%2 == 0 {
			n, _ := strconv.Atoi(string(ch))
			for j := 0; j < n; j++ {
				new = append(new, strconv.Itoa(count))
			}
			count++
		} else {
			n, _ := strconv.Atoi(string(ch))
			for j := 0; j < n; j++ {
				new = append(new, ".")
			}
		}
	}
	// fmt.Println(new)

	for i := 0; i < len(new); i++ {
		if new[i] == "." {
			for j := len(new) - 1; j > i; j-- {
				if new[j] != "." {
					part1 := append(new[:i], new[j])
					part2 := append(part1, new[i+1:j]...)
					part3 := append(part2, ".")
					part4 := append(part3, new[j+1:]...)
					new = part4
					break
				}
			}
		}
	}

	// fmt.Println(new)

	fin := 0
	for i, ch := range new {
		if ch != "." {
			n, _ := strconv.Atoi(string(ch))
			fin += i * n
		}
	}

	fmt.Println(fin)
}
