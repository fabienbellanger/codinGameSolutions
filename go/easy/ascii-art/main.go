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
	text := scanner.Text()

	// Alphabet en ASCII Art
	letters := make([27][L * H]string)
	
    for i := 0; i < H; i++ {
        scanner.Scan()
        row := scanner.Text()
    }
    
    // fmt.Fprintln(os.Stderr, "Debug messages...")
    fmt.Println("answer")// Write answer to stdout
}
