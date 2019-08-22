package main

import (
    "fmt"
    "os"
)

/*
Solution test
-------------
144
156
12
24
288
12
*/

func main() {
    var N int
    fmt.Scan(&N)

    results := make(map[int]int)
    fmt.Fprintf(os.Stderr, "%v - %t\n", results, isMapFull(results, N))
    
    for i := 0; i < N; i++ {
        var operation, arg1, arg2 string
        fmt.Scan(&operation, &arg1, &arg2)
    }

    for i := 0; i < N; i++ {
        
        // fmt.Fprintln(os.Stderr, "Debug messages...")
        fmt.Println("1")
    }
}

func isMapFull(results map[int]int, size int) bool {
    return len(results) == size
}
