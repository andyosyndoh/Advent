package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("instructions.txt")

	content := strings.Split(string(input), " ")

	count := 0
	for count < 25 {
		count++
		i := 0
		for i < len(content) {
			if content[i] == "0" {
				content[i] = "1"
				i++
			} else if len(content[i]) == 1 || len(content[i])%2 != 0 {
				num, _ := strconv.Atoi(content[i])
				content[i] = strconv.Itoa(num*2024)
				i++
			} else  if len(content[i])%2 == 0 {
				last := new(content[i+1:])
				part1,part2 := content[i][:len(content[i])/2],content[i][len(content[i])/2:]
				p2 , _ := strconv.Atoi(part2)
				part2 = strconv.Itoa(p2)
				cont1 := append(content[:i], []string{part1,part2}...)
				cont2 := append(cont1, last...)
				content = cont2
				i += 2
			}
		}
	}
	fmt.Println(len(content))
}

func new(s []string) []string {
	n := []string{}
	for _ , v := range s {
		n = append(n, v)
	}
	return n
}