package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
    input, _ := os.ReadFile("instructions.txt")
    content := strings.Fields(string(input))

    for count := 0; count < 75; count++ {
        fmt.Println(count)
        newContent := make([]string, 0, len(content)*2)

        for _, item := range content {
            switch {
            case item == "0":
                newContent = append(newContent, "1")
            case len(item) == 1 || len(item)%2 != 0:
                num, _ := strconv.Atoi(item)
                newContent = append(newContent, strconv.Itoa(num*2024))
            default:
                mid := len(item) / 2
                part1, part2 := item[:mid], item[mid:]
                p2, _ := strconv.Atoi(part2)
                newContent = append(newContent, part1, strconv.Itoa(p2))
            }
        }
        content = newContent
    }
    fmt.Println(len(content))
}
