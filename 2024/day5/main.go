package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("instructions.txt")

	lines := strings.Split(string(input), "\n")

	pages := make(map[string][]string)
	var sequence [][]string

	for _, line := range lines {
		if line != "" {
			if strings.Contains(line, "|") {
				parts := strings.Split(line, "|")
				pages[parts[0]] = append(pages[parts[0]], parts[1])
			} else {
				parts := strings.Split(line, ",")
				sequence = append(sequence, parts)
			}
		}
	}

	// fmt.Println(pages)

	num := 0
	var newsequence [][]string
	for _, seq := range sequence {
		n := check(seq, pages)
		if n != -1 {
			num += n
		} else {
			newsequence = append(newsequence, seq)
		}
	}

	res := 0
	for _, seq := range newsequence {
		res += resolve(seq, pages)
	}

	fmt.Println(num)
	fmt.Println(res)
}

func check(seq []string, pages map[string][]string) int {
	for i, page := range seq {
		if cor, ok := pages[page]; ok {
			for j := 0; j < i; j++ {
				for _, n := range cor {
					if seq[j] == n {
						return -1
					}
				}
			}
		}
	}

	index := len(seq) / 2
	n, _ := strconv.Atoi(seq[index])
	return n
}

func resolve(seq []string, pages map[string][]string) int {
	if solve(seq, pages) {
		// fmt.Println(seq)
		index := len(seq) / 2
		n, _ := strconv.Atoi(seq[index])
		return n
	}
	index := len(seq) / 2
	n, _ := strconv.Atoi(seq[index])
	return n
}

func solve(seq []string, pages map[string][]string) bool {
	// fmt.Println(seq)
	for i := 0; i < len(seq); i++ {
		if cor, ok := pages[seq[i]]; ok {
			for j := 0; j < i; j++ {
				for _, n := range cor {
					if seq[j] == n {
						part1 := append(seq[:j], seq[j+1:i+1]...)
						// fmt.Println(part1)
						part2 := append([]string{n}, seq[i+1:]...)
						// fmt.Println(part2)
						part1 = append(part1, part2...)
						// fmt.Println(part1)
						if solve(part1, pages) {
							return true
						}
					}
				}
			}
		}
	}
	return true
}
