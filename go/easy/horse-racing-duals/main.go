package main

import (
	"fmt"
)

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

/*

OPTI
----

m := make(map[int]int, 0)
pi...
if gap <len(m) {
	loop de pi - gap à pi + gap
} else {
	loop sur m
}

*/


func main() {
    var n int
    fmt.Scan(&n)
	
	gap := 10000000
	list := make([]int, 0, n)
	stop := false

    for i := 0; i < n; i++ {
        var pi int
		fmt.Scan(&pi)
		
		if !stop {
			for j, v := range list {
				var g int
				// fmt.Printf("%d - %d\n", i, j)
				if pi >= v {
					g = pi - v
				} else {
					g = v - pi
				}

				if g < gap {
					gap = g

					if gap == 1 {
						stop = true
						break
					}
				}
			}

			list = append(list, pi)
		}
	}
    
	// Output
	// ------
    fmt.Println(gap)
}
