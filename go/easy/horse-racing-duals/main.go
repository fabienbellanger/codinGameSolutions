package main

import (
	"fmt"
)

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
    var n int
    fmt.Scan(&n)
	
	gap := 10000000
	list := make([]int, 0)
	stop := false

    for i := 0; i < n; i++ {
        var pi int
		fmt.Scan(&pi)
		
		if !stop {
			for _, v := range list {
				var g int

				if pi >= v {
					g = pi - v
				} else {
					g = v - pi
				}

				if g < gap {
					gap = g
					break
				}
			}

			if gap == 1 {
				stop = true
			}

			list = append(list, pi)
		}
	}
    
	// Output
	// ------
    fmt.Println(gap) // Write answer to stdout
}
