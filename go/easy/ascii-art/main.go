package main

import (
	"fmt"
	"os"
	"bufio"
)

type lineType []string

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
	// text := scanner.Text()

	// Alphabet en ASCII Art
	// line := make([]string, L * H)
	// var letters [27]lineType
	
    for i := 0; i < H; i++ {
        scanner.Scan()
		row := scanner.Text()
		

		println("ROWS", len(row))

		/*
			k = -1
			l = 0
			loop 0 <= j < 27 * L
				if j % L == 0
					k++
					l = 0

				letters[k][i * N + l] = row[j]

				l++
		*/
    }
    
    // fmt.Fprintln(os.Stderr, "Debug messages...")
    fmt.Println("answer")// Write answer to stdout
}
