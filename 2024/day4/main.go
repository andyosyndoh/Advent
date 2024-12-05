package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("instructions.txt")

	split := strings.Split(string(input), "\n")

	count := 0

	count = Verticalbackcount(split, count)
	count = CountDown(split, count)
	count = Diagonalyfront(split, count)
	count = Diagonalyback(split, count)
	fmt.Println(count)

	xmas := countxmas(split)

	fmt.Println(xmas)
}

func Verticalbackcount(lines []string, count int) int {
	str1 := "SAMX"
	str2 := "XMAS"
	for _, line := range lines {
		for i := 0; i < len(line); i++ {
			if i+len(str1)-1 < len(line) {
				if line[i:i+len(str1)] == str1 || line[i:i+len(str1)] == str2 {
					count++
				}
			}
		}
	}
	return count
}

func CountDown(lines []string, count int) int {
	str1 := "XMAS"
	str2 := "SAMX"
	for i, line := range lines {
		for j := range line {
			if i <= len(lines)-len(str1) {
				s := string(lines[i][j]) + string(lines[i+1][j]) + string(lines[i+2][j]) + string(lines[i+3][j])
				if s == str1 || s == str2 {
					count++
				}
			}
		}
	}
	return count
}

func Diagonalyfront(lines []string, count int) int {
	str1 := "XMAS"
	str2 := "SAMX"
	for i, line := range lines {
		for j := range line {
			if (i <= len(lines)-len(str1)) && (j <= len(line)-len(str1)) {
				s := string(lines[i][j]) + string(lines[i+1][j+1]) + string(lines[i+2][j+2]) + string(lines[i+3][j+3])
				// fmt.Println(s)
				if s == str1 || s == str2 {
					count++
				}
			}
		}
	}
	return count
}

func Diagonalyback(lines []string, count int) int {
	str1 := "XMAS"
	str2 := "SAMX"
	for i, line := range lines {
		for j := range line {
			if (i <= len(lines)-len(str1)) && (j >= len(str1)-1) {
				s := string(lines[i][j]) + string(lines[i+1][j-1]) + string(lines[i+2][j-2]) + string(lines[i+3][j-3])
				// fmt.Println(s,i,j)
				if s == str1 || s == str2 {
					count++
				}
			}
		}
	}
	return count
}

func countxmas(lines []string) int {
	count := 0
	str := []string{"MMASS", "MSAMS", "SSAMM", "SMASM"}
	for i, line := range lines {
		for j := range line {
			if i <= len(line) - 3 && j <= len(line) -3 {
				s := string(lines[i][j]) + string(lines[i][j+2]) + string(lines[i+1][j+1]) + string(lines[i+2][j]) + string(lines[i+2][j+2])
				for _, v := range str {
					if s == v {
                        count++
                    }
				}
			}
		}
	}
	return count
}
