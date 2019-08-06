package main

import "fmt"
import "os"
import "bufio"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Buffer(make([]byte, 1000000), 1000000)

    var n int
    scanner.Scan()
    fmt.Sscan(scanner.Text(), &n)
    
    for i := 0; i < n; i++ {
        scanner.Scan()
        isbn := scanner.Text()
        // isbn := scanner.Bytes()
        ok := checkISBN(isbn)
        println(ok)
    }
    
    // fmt.Fprintln(os.Stderr, "Debug messages...")
    fmt.Println("answer") // Write answer to stdout
}

// checkISBN checks if ISBN is correct
func checkISBN(s string) bool {
    if len(s) != 10 && len(s) != 13 {
        return false
    }

    return true
}
