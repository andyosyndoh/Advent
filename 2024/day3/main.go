package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("getinput.txt")

	regex := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`)

	matches := regex.FindAllString(string(input), -1)
	ans := 0
	do := true
	for _ , str := range matches{
		// fmt.Println(str)
		if str == "do()" {
			do = true
		}
		if str == "don't()" {
			do = false
		}
		if do && str != "do()" && str != "don't()"{
			str = strings.TrimLeft(str, "mul(" )
			str = strings.TrimRight(str, ")")
			parts := strings.Split(str, ",")
			f1, _ := strconv.Atoi(parts[0])
			f2, _ := strconv.Atoi(parts[1])
			ans += f1*f2
		}
	}
	fmt.Println(ans)
}
