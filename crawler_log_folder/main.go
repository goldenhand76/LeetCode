package main

import "fmt"

func main() {
	r := minOperations([]string{"d1/", "d2/", "../", "d21/", "./"})
	fmt.Print(r)
}

func minOperations(logs []string) int {
	main := 0
	for _, v := range logs {
		switch v {
		case "./":

		case "../":
			if main > 0 {
				main--
			}

		default:
			main++
		}
	}
	return main
}
