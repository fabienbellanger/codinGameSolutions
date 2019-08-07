package main

import (
	"fmt"
)

func main() {
    var n int
    fmt.Scan(&n)
	
	gap := 10000000
	list := make(map[int]int, 0)
	stop := false

    for i := 0; i < n; i++ {
        var pi int
		fmt.Scan(&pi)
		
		if !stop {
			if gap < len(list) {
				for i := pi - gap; i < pi + gap && !stop; i++ {
					if _, ok := list[i]; ok {
						updateGap(pi, i, &gap, &stop)
					}
				}
			} else {
				for k := range list {
					updateGap(pi, k, &gap, &stop)

					if stop {
						break
					}
				}
			}

			list[pi] = 0
		}
	}
    
	// Output
	// ------
    fmt.Println(gap)
}

func updateGap(pi int, i int, gap *int, stop *bool) {
	var g int

	if pi >= i {
		g = pi - i
	} else {
		g = i - pi
	}
	
	if g < *gap {
		*gap = g

		if *gap == 1 {
			*stop = true
		}
	}
}
