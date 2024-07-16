package main

import (
	"fmt"
	"sort"
	"strconv"
	"unicode"
)

func main() {
	result := countOfAtoms("K4(ON(SO3)2)2")
	fmt.Print(result)
}

func countOfAtoms(formula string) string {
	f := []rune(formula)
	letters := []string{}
	numbers := []int{}
	p1 := 0
	p2 := 0
	p3 := 0
	p4 := 0
	for p2 < len(f) {
		if unicode.IsUpper(f[p2]) {
			p4 = 1
			if p2+1 < len(f) && unicode.IsDigit(f[p2+p4]) {
				for p2+p4 < len(f) && unicode.IsDigit(f[p2+p4]) {
					p4++
					continue
				}
				p4 -= 1
				num, _ := strconv.Atoi(string(f[p2+1 : p2+p4+1]))
				letters = append(letters, string(f[p2]))
				numbers = append(numbers, num)

			} else if p2+1 < len(f) && unicode.IsLower(f[p2+1]) {
				letters = append(letters, string(f[p2:p2+2]))
				if p2+2 < len(f) && unicode.IsDigit(f[p2+p4+1]) {
					for p2+p4+1 < len(f) && unicode.IsDigit(f[p2+p4+1]) {
						p4++
						continue
					}
					p4 -= 1
					num, _ := strconv.Atoi(string(f[p2+2 : p2+p4+2]))
					numbers = append(numbers, num)
				} else {
					numbers = append(numbers, 1)
				}

			} else {
				letters = append(letters, string(f[p2]))
				numbers = append(numbers, 1)
			}
		} else if p2+1 < len(f) && string(f[p2]) == ")" && unicode.IsDigit(f[p2+1]) {
			p4 = 1
			for p2+p4 < len(f) && unicode.IsDigit(f[p2+p4]) {
				p4++
				continue
			}
			p4 -= 1
			m1, _ := strconv.Atoi(string(f[p2+1 : p2+p4+1]))
			p1 = p2 - 1
			p3 = len(numbers) - 1
			for string(f[p1]) != "(" {
				if unicode.IsUpper(f[p1]) {
					numbers[p3] = numbers[p3] * m1
					p3--
					p1--
				} else if unicode.IsLower(f[p1]) {
					numbers[p3] = numbers[p3] * m1
					p3--
					p1 = p1 - 2
				} else {
					p1--
				}
			}
			f = append(f[:p1], f[p1+1:]...)
		}
		p2++
	}

	// Old Solution

	elementCounts := map[string]int{}
	for i := 0; i < len(letters); i++ {
		elementCounts[letters[i]] += numbers[i]
	}

	// Sort elements and construct the resulting string
	keys := make([]string, 0, len(elementCounts))
	for k := range elementCounts {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	result := ""
	for _, k := range keys {
		result += k
		if elementCounts[k] > 1 {
			result += strconv.Itoa(elementCounts[k])
		}
	}

	return result
}
