package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Buffer(make([]byte, 1000000), 1000000)

    var L int
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&L)
    
    var H int
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&H)
    
    scanner.Scan()
    //T := scanner.Text()
    for i := 0; i < H; i++ {
        scanner.Scan()
        //ROW := scanner.Text()
    }
    
    // fmt.Fprintln(os.Stderr, "Debug messages...")
    fmt.Println("answer")// Write answer to stdout
}
