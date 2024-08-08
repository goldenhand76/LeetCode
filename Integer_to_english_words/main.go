package main

import (
	"fmt"
	"runtime/debug"
	"strings"
)

type Digits struct {
	firstDigit  []string
	secondDigit []string
}

func numberToWords(num int) string {
	digits := Digits{
		firstDigit: []string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"},
		secondDigit: []string{"Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen",
			"Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"},
	}
	answer := []string{}
	digits.nameInteger(num, &answer)
	return strings.Join(answer, " ")
}

func (d Digits) nameInteger(num int, ans *[]string) {
	if num == 0 {
		*ans = append(*ans, "Zero")
		return
	}
	if num < 10 {
		*ans = append(*ans, d.firstDigit[num])

	} else if num < 20 {
		units := num % 10
		*ans = append(*ans, d.secondDigit[units])

	} else if num >= 20 && num <= 99 {
		tens := num / 10
		units := num % 10
		*ans = append(*ans, d.secondDigit[tens+8])
		if units != 0 {
			*ans = append(*ans, d.firstDigit[units])
		}

	} else if num > 99 && num <= 999 {
		hundreds := num / 100
		*ans = append(*ans, d.firstDigit[hundreds])
		*ans = append(*ans, "Hundred")

		units := num % 100
		if units != 0 {
			d.nameInteger(units, ans)
		}

	} else if num <= 999999 {
		thousands := num / 1000
		d.nameInteger(thousands, ans)
		*ans = append(*ans, "Thousand")

		hundreds := num % 1000
		if hundreds != 0 {
			d.nameInteger(hundreds, ans)
		}

	} else if num <= 999999999 {
		millions := num / 1000000
		d.nameInteger(millions, ans)
		*ans = append(*ans, "Million")

		thousands := (num % 1000000) / 1000
		if thousands != 0 {
			d.nameInteger(thousands, ans)
			*ans = append(*ans, "Thousand")
		}

		hundreds := num % 1000
		if hundreds != 0 {
			d.nameInteger(hundreds, ans)
		}
	} else if num <= 999999999999 {
		billions := num / 1000000000
		d.nameInteger(billions, ans)
		*ans = append(*ans, "Billion")

		millions := num % 1000000000 / 1000000
		if millions != 0 {
			d.nameInteger(millions, ans)
			*ans = append(*ans, "Million")
		}

		thousands := (num % 1000000) / 1000
		if thousands != 0 {
			d.nameInteger(thousands, ans)
			*ans = append(*ans, "Thousand")
		}

		hundreds := num % 1000
		if hundreds != 0 {
			d.nameInteger(hundreds, ans)
		}
	}
}

func init() {
	debug.SetGCPercent(-1)
	debug.FreeOSMemory()
}
func main() {
	fmt.Println(numberToWords(100))
	fmt.Println(numberToWords(1000000))
	fmt.Println(numberToWords(1000))
}
