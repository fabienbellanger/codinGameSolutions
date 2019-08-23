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

	printGrid(grid)
	zeros := getZeroNumbers(grid)

	fmt.Println(zeros)
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
