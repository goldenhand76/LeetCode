package main

import (
	"fmt"
)

func main() {
	r := maximumGain("aabbaaxybbaabb", 5, 4)
	fmt.Print(r)
}

func maximumGain(s string, x int, y int) int {
	o := []rune(s)
	p1 := 0
	points := 0
	if x > y {
		for p1 < len(o)-1 {
			if string(o[p1:p1+2]) == "ab" {
				o = append(o[:p1], o[p1+2:]...)
				points += x
				p1 -= 2
			}
			if p1 < 0 {
				p1 = 0
			} else {
				p1++
			}
		}
		p1 = 0
		for p1 < len(o)-1 {
			if string(o[p1:p1+2]) == "ba" {
				o = append(o[:p1], o[p1+2:]...)
				points += y
				p1 -= 2
			}
			if p1 < 0 {
				p1 = 0
			} else {
				p1++
			}
		}
	} else {
		for p1 < len(o)-1 {
			if string(o[p1:p1+2]) == "ba" {
				o = append(o[:p1], o[p1+2:]...)
				points += y
				p1 -= 2
			}
			if p1 < 0 {
				p1 = 0
			} else {
				p1++
			}
		}
		p1 = 0
		for p1 < len(o)-1 {
			if string(o[p1:p1+2]) == "ab" {
				o = append(o[:p1], o[p1+2:]...)
				points += x
				p1 -= 2
			}
			if p1 < 0 {
				p1 = 0
			} else {
				p1++
			}
		}
	}
	return points
}
