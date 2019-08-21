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

    results := make([]int, N)
    fmt.Fprintf(os.Stderr, "%v\n", results)
    
    for i := 0; i < N; i++ {
        var operation, arg1, arg2 string
        fmt.Scan(&operation, &arg1, &arg2)
    }

    for i := 0; i < N; i++ {
        
        // fmt.Fprintln(os.Stderr, "Debug messages...")
        fmt.Println("1")// Write answer to stdout
    }
}
