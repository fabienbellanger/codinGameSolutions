package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	const GAP = 3

	var W, H int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &W, &H)

	grid := make([]string, H, H)

	for i := 0; i < H; i++ {
		scanner.Scan()
		grid[i] = scanner.Text()
	}

	// Parcours des lignes
	// -------------------
	for i := 0; i < W; i += GAP {
		k := i
		response := string(grid[0][i])

		for j := 1; j < H-1; j++ {
			// On regarde à droite
			if i < W-1 && grid[j][i+1] == '-' {
				i += GAP
			} else if i-1 > 0 && grid[j][i-1] == '-' {
				i -= GAP
			}
		}

		response += string(grid[H-1][i])
		fmt.Println(response)

		i = k
	}
}
