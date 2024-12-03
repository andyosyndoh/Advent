package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("instructions.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")
	var first []int
	var second []int
	for _, line := range lines {
		parts := strings.Split(line, "   ")
		f, err1 := strconv.Atoi(parts[0])
		s, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil{
			panic(err1)
		}
		first = append(first, f)
		second = append(second, s)
	}

	sort.Ints(first)
	sort.Ints(second)

	ans := 0
	for i, num := range first {
		num2 := second[i]
		dif := num2 - num
		if dif < 0 {
			dif = -dif
		}
		ans += dif
	}

	fmt.Println(ans)

	diff := 0
	for _, num := range first {
		count := 0
		for _, num2 := range second {
			if num == num2 {
				count++
			}
		}
		num *= count
		diff += num
	}

	fmt.Println(diff)
}
