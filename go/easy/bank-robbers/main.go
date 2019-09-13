package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var R int
	fmt.Scan(&R)

	var V int
	fmt.Scan(&V)

	// Exemple
	// -------
	// C = 3 et N = 1
	// => {0-9} {A E I O U} {A E I O U}
	// Test1 : 10 * 5 * 5 = 250
	// Test2 = 1125

	robbers := make([]int, R)
	robber := 0
	for i := 0; i < V; i++ {
		var C, N int
		fmt.Scan(&C, &N)

		robbers[robber] += int(math.Pow10(N) * math.Pow(5.0, float64(C-N)))
		fmt.Fprintln(os.Stderr, "Robbers: %v - Robber: %d\n", robbers, robber)

		robber = (robber + 1) % R
	}

	// Recherche du max robber time
	// TODO
	response := 0

	fmt.Println(response)
}
