package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("instructions.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	lines := strings.Split(string(content), "\n")

	fmt.Println(len(lines))
	safe := 0
	for j := 0; j < len(lines); j++ {
		slope := 0
		line := lines[j]
		// fmt.Println(line)
		parts := strings.Split(line, " ")
		// fmt.Println(parts)
		num1, err1 := strconv.Atoi(parts[0])
		num2, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			fmt.Println(err1)
			os.Exit(1)
		}
		if num2 > num1 {
			slope = 1
		} else if num2 < num1 {
			slope = -1
		}

		
		if check(parts, slope) {
			safe += 1
		} else {
			if check2(parts) {
				safe += 1
			}
		}
		// fmt.Println()
	}
	fmt.Println((safe))
}

func check(parts []string, slope int) bool {
	if slope == 0 {
		return false
	}
	for i := 1; i < len(parts); i++ {
		curr, err1 := strconv.Atoi(parts[i])
		prevv, err2 := strconv.Atoi(parts[i-1])
		if err1 != nil || err2 != nil {
			fmt.Println(err1)
			os.Exit(1)
		}
		if curr > prevv && slope == -1 {
			// fmt.Println(curr, prevv, "inc")
			return false
		}
		if curr < prevv && slope == 1 {
			// fmt.Println(curr, prevv, "dec")
			return false
		}
		diff := curr - prevv
		if diff > 3 || diff < -3 || diff == 0 {
			// fmt.Println(curr, prevv, "pro")
			return false
		}
	}
	return true
}

func check2(parts []string) bool {
	for i := range parts {
		new := []string{}
		for j, d :=range parts {
			if i != j {
				new = append(new, d)
			}
		}
		num1, _ := strconv.Atoi(new[0])
		num2, _ := strconv.Atoi(new[1])

		slope := 0
		if num2 > num1 {
			slope = 1
		} else if num2 < num1 {
			slope = -1
		}
		if check(new, slope) {
			return true
		}
	}
	return false
}