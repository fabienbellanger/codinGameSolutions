package main

import (
	"fmt"
)

func main() {
    var width, height int
    fmt.Scan(&width, &height)
    
    for i := 0; i < height; i++ {
        var line string
        fmt.Scan(&line)
    }
    var side string
    fmt.Scan(&side)
    
    for i := 0; i < height; i++ {
        
        // fmt.Fprintln(os.Stderr, "Debug messages...")
        fmt.Println("#####")// Write action to stdout
    }
}
