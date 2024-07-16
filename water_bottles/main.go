package main

import "fmt"

func main() {
	fmt.Println(numWaterBottles(9, 3))
}

func numWaterBottles(numBottles int, numExchange int) int {
	newBottles := numBottles
	for newBottles >= numExchange {
		numBottles += (newBottles / numExchange)
		newBottles = (newBottles / numExchange) + (newBottles % numExchange)

	}
	return numBottles
}
