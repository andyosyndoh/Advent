package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("dimensions.txt")
	if err != nil {
		fmt.Println("error")
		return
	}

	boxes := strings.Split(string(file), "\n")
	res := 0
	for _, box := range boxes {
		if box != "" {
			sides := strings.Split(box, "x")
			w, _ := strconv.Atoi(strings.TrimSpace(sides[0]))
			h, _ := strconv.Atoi(strings.TrimSpace(sides[1]))
			l, _ := strconv.Atoi(strings.TrimSpace(sides[2]))
			tot := 2*l*w + 2*w*h + 2*h*l
			bow := l*w*h
			small := []int{}
			small = append(small, w)
			small = append(small, h)
			small = append(small, l)
			sort.Ints(small)
			s := w*l
			if w*h < s {
				s = w*h
			}
			if h*l < s {
				s = h*l
			}
			tot += s
			res += bow + (small[0]+small[0]+small[1]+small[1])
		}
	}
	fmt.Println(res)
}
