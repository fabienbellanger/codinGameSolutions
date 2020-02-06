package main

import "fmt"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	var n, x, y int
	fmt.Scan(&n)

	/*
	   x, y > n
	   n < x <= 2*n
	   y = (n*x) / x - n

	   x = 2*n
	   tant que x > n
	       x--
	       y = (n*x) / (x-n)
	       if (n*x) % (x-n) == 0 and y > n alors
	           found solution: y = (n*x) / (x-n)
	       sinon KO
	*/

	for x = n + 1; x <= 2*n; x++ {
		if (n*x)%(x-n) == 0 {
			y = int((n * x) / (x - n))
			fmt.Printf("1/%d = 1/%v + 1/%v\n", n, y, x)
		}
	}
}
