package main

import "fmt"

func main() {
	r := averageWaitingTime([][]int{{5, 2}, {5, 4}, {10, 3}, {20, 1}})
	fmt.Print(r)
}

func averageWaitingTime(customers [][]int) float64 {
	turnAroundTime := 0
	serviceTime := 0
	total := float64(len(customers))
	for _, customer := range customers {
		if customer[0] >= serviceTime {
			serviceTime = customer[0] + customer[1]
			turnAroundTime += serviceTime - customer[0]
		} else {
			serviceTime += customer[1]
			turnAroundTime += serviceTime - customer[0]
		}

	}
	return float64(turnAroundTime) / total

}
