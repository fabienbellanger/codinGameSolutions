package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	var N int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &N)

	var L int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &L)

	grid := make([][]int, N, N)

	for i := 0; i < N; i++ {
		scanner.Scan()
		line := scanner.Text()
		tab := strings.Split(line, " ")

		for _, v := range tab {
			if v == "C" {
				grid[i] = append(grid[i], L)
			} else {
				grid[i] = append(grid[i], 0)
			}
		}
	}

	for y := 0; y < N; y++ {
		for x := 0; x < N; x++ {
			if grid[y][x] == L {
				for k := 1; k < L; k++ {
					// First line
					// ----------
					for l := x - k; l <= x+k; l++ {
						if y-k >= 0 && l >= 0 && l < N && grid[y-k][l] < L-k {
							grid[y-k][l] = L - k
						}
					}

					// Second line
					// -----------
					for l := y - k + 1; l <= y+k; l++ {
						if x+k < N && l >= 0 && l < N && grid[l][x+k] < L-k {
							grid[l][x+k] = L - k
						}
					}

					// Third line
					// ----------
					for l := x - k; l <= x+k-1; l++ {
						if y+k < N && l >= 0 && l < N && grid[y+k][l] < L-k {
							grid[y+k][l] = L - k
						}
					}

					// Last line
					// ---------
					for l := y - k + 1; l <= y+k-1; l++ {
						if x-k >= 0 && l >= 0 && l < N && grid[l][x-k] < L-k {
							grid[l][x-k] = L - k
						}
					}
				}
			}
		}
	}

	// printGrid(grid)

	fmt.Println(getZeroNumbers(grid))
}

func printGrid(g [][]int) {
	for _, line := range g {
		fmt.Fprintln(os.Stderr, line)
	}
}

func getZeroNumbers(g [][]int) int {
	n := 0

	for _, line := range g {
		for _, cell := range line {
			if cell == 0 {
				n++
			}
		}
	}

	return n
}
