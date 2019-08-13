package main

import (
	"fmt"
	"strconv"
)

func main() {
    var width, height int
    fmt.Scan(&width, &height)
	
	grid := make([]string, 0)
    for i := 0; i < height; i++ {
        var line string
		fmt.Scan(&line)
		grid = append(grid, line)
	}

    for i := 0; i < height; i++ {
		line := ""

		for j, c := range grid[i] {
			if c == '0' {
				n := 0

				if i + 1 < height && grid[i + 1][j] == '0' {
					n++
				}
				if i - 1 >= 0 && grid[i-1][j] == '0' {
					n++
				}
				if j + 1 < width && grid[i][j+1] == '0' {
					n++
				}
				if j - 1 >= 0 && grid[i][j-1] == '0' {
					n++
				}

				line += strconv.Itoa(n)
			} else {
				line += "#"
			}
		}

		fmt.Println(line)
    }
}
