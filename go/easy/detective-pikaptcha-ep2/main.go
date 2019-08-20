package main

import (
    "fmt"
    "os"
)

func main() {
    var width, height int
    fmt.Scan(&width, &height)
    
    grid := make([][]byte, 0)
    startX := 0
    startY := 0
    for i := 0; i < height; i++ {
        var line string
		fmt.Scan(&line)
		grid = append(grid, []byte(line))
    }

    // Position de départ
    // ------------------
    startFound := false
    for i := 0; i < height && !startFound; i++ {
        for j := 0; j < width && !startFound; j++ {
            c := grid[i][j]

            if c == '>' || c == '<' || c == 'v' || c == '^' {
                startX = j
                startY = i
                startFound = true
            }
        }
    }
    
    var side string
    fmt.Scan(&side)

    fmt.Fprintf(os.Stderr, "        Start X: %d\n", startX)
    fmt.Fprintf(os.Stderr, "        Start Y: %d\n", startY)
    fmt.Fprintf(os.Stderr, "Start direction: %s\n", string(grid[startX][startY]))
    fmt.Fprintf(os.Stderr, "      Wall side: %s\n\n", side)
    
    for i := 0; i < height; i++ {
        
        // fmt.Fprintln(os.Stderr, "Debug messages...")
        fmt.Println("#####")// Write action to stdout
    }
}
