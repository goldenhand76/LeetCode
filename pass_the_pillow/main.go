package main

import "fmt"

func main() {
	fmt.Println(passThePillow(4, 5))
	fmt.Println(passThePillow(3, 2))
}

func passThePillow(n int, time int) int {
	step := n - 1
	if (time/step)%2 == 0 {
		return 1 + (time % step)
	} else {
		return n - (time % step)
	}
}
