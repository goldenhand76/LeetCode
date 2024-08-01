package main

import (
	"strconv"
)

func main() {
	countSeniors([]string{"12345678911321"})
}

func countSeniors(details []string) int {
	count := 0
	for i := range details {
		age, _ := strconv.Atoi(string([]rune(details[i])[11]))
		if age >= 6 {
			if age == 6 {
				age2, _ := strconv.Atoi(string([]rune(details[i])[12]))
				if age2 == 0 {
					break
				}
			}
			count++
		}

	}
	return count
}
