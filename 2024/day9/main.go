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
	fmt.Println(new)

	cop := form(new)

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

	try := make(map[string]bool)

	for i := len(cop)-1; i >= 0; i-- {
		if cop[i] != "." {
			tempnum, con := getnums(cop ,i)
			if len(tempnum) > 0 {
				if try[tempnum[0]] {
					continue
				}
				try[tempnum[0]] = true
			}
			for j := 0; j < con; j++ {
				if cop[j] == "." {
					tempdots, end := getdots(cop, j)
					if len(tempdots) >= len(tempnum) {
						dif := len(tempdots) - len(tempnum)
						if dif > 0 {
							for k := 0;k < dif; k++ {
								tempnum = append(tempnum, ".")
							}
						}
						// fmt.Println(tempdots)
						// fmt.Println(tempnum)
						part1 := append(cop[:j], tempnum...)
						part2 := append(part1, cop[end:con+1]...)
						part3 := append(part2, tempdots[:len(tempdots)-dif]...)
						part4 := append(part3, cop[i+1:]...)
						cop = part4
						// fmt.Println(cop)
						i = con+1
						break
					} else {
						j = end
					}
				}
			}
		}
	}
 
	fmt.Println(cop)

	fin := 0
	for i, ch := range new {
		if ch != "." {
			n, _ := strconv.Atoi(string(ch))
			fin += i * n
		}
	}

	fmt.Println(fin)
	fin = 0

	for i, ch := range cop {
		if ch != "." {
			n, _ := strconv.Atoi(string(ch))
			fin += i * n
		}
	}

	fmt.Println(fin)
}

func form(new []string) []string {
	cop := []string{}
	cop = append(cop, new...)
	return cop
}

func getnums(cop []string, i int) ([]string, int) {
	for j := i; j >= 0; j-- {
		if cop[j] != cop[i] {
			return form(cop[j+1:i+1]), j
		}
	}
	return []string{}, 0
}

func getdots(cop []string, i int) ([]string, int) {
	for j := i; j < len(cop); j++ {
		if cop[j] != cop[i] {
			return form(cop[i:j]), j
		}
	}
	return []string{}, 0
}