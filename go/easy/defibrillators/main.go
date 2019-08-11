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

    var LON string
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&LON)
    
    var LAT string
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&LAT)
    
    var N int
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&N)
    
    for i := 0; i < N; i++ {
        scanner.Scan()
        //DEFIB := scanner.Text()
    }
    
    // fmt.Fprintln(os.Stderr, "Debug messages...")
    fmt.Println("answer")// Write answer to stdout
}
