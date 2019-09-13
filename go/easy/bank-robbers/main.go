package main

import (
	"fmt"
	"math"
)

func main() {
	var R int
	fmt.Scan(&R)

	var V int
	fmt.Scan(&V)

	robbers := make([]int, R)
	robber := 0
	for i := 0; i < V; i++ {
		var C, N int
		fmt.Scan(&C, &N)

		robbers[robber] += int(math.Pow10(N) * math.Pow(5.0, float64(C-N)))

		robber = getIndexMinTime(robbers)
	}

	// Recherche du max robber time
	response := 0
	for _, v := range robbers {
		if v > response {
			response = v
		}
	}

	fmt.Println(response)
}

func getIndexMinTime(g []int) int {
	min := math.MaxInt64
	index := 0

	for i, v := range g {
		if v < min {
			min = v
			index = i
		}
	}

	return index
}
