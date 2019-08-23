package main

import (
	"fmt"
	"os"
)

func main() {
	var expression string
	fmt.Scan(&expression)

	chars := make(map[byte]bool)
	fmt.Fprintf(os.Stderr, "%v\n", chars)

	// fmt.Fprintln(os.Stderr, "Debug messages...")
	fmt.Println("true/false") // Write answer to stdout
}
