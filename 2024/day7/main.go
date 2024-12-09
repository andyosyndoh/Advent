package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("instructions.txt")

	content := strings.Split(string(input), "\n")

	count := 0
	for _, line := range content {
		parts := strings.Split(line, ":")
		if solve(parts[0], parts[1]) {
			// count++
			s, _ := strconv.Atoi(parts[0])
			count += s
		}
	}

	fmt.Println(count)
}

func solve(ans, nums string) bool {
	nums = strings.TrimSpace(nums)
	n := strings.Split(nums, " ")
	op := GeneratePermutations(len(n) - 1)
	// fmt.Println(n, op, ans)
	for _, op1 := range op {
		s := strings.Split(op1, " ")
		if check(n, s, ans) {
			return true
		}
	}
	return false
}

func check(n, ops []string, ans string) bool {
	str := []string{n[0]}
	for i := 1; i < len(n); i++ {
		str = append(str, n[i])
		str = append(str, ops[i-1])
	}
	// fmt.Println(str)
	nums := []int{}
	for _, ch := range str {
		if ch == "+" || ch == "*" || ch == "||"{
			a := nums[len(nums)-2]
			b := nums[len(nums)-1]
			nums = nums[:len(nums)-2]

			calc := 0
			if ch == "+" {
				calc = a + b
			}
			if ch == "*" {
				calc = a * b
			}
			if ch == "||" {
				as := strconv.Itoa(a)
				bs := strconv.Itoa(b)
				f := as+bs
				calc , _ = strconv.Atoi(f) 
			}
			nums = append(nums, calc)
		} else {
			s, _ := strconv.Atoi(ch)
			nums = append(nums, s)
		}
	}
	// fmt.Println(nums[0])
	aa, _ := strconv.Atoi(ans)
	return nums[0] == aa
}

func GeneratePermutations(slots int) []string {
	if slots <= 0 {
		return []string{}
	}
	
	var results []string
	var generate func(current string, remaining int)

	// Recursive function to build permutations
	generate = func(current string, remaining int) {
		if remaining == 0 {
			results = append(results, current)
			return
		}

		// Add '*' and recurse
		generate(current+"* ", remaining-1)
		// Add '+' and recurse
		generate(current+"+ ", remaining-1)

		generate(current+"|| ", remaining-1)
	}

	// Start with an empty string and the full number of slots
	generate("", slots)
	return results
}