package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		var A, B int
		fmt.Scan(&A, &B)
	}

	// fmt.Fprintln(os.Stderr, "Debug messages...")
	fmt.Println("42")
}
