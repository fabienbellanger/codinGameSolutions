package main

import (
	"fmt"
	"time"
	"os"
)

func main() {
    var n int
	fmt.Scan(&n)
	
	start := time.Now()
	
	gap := 10000000
	list := make(map[int]int, 0)
	stop := false

    for i := 0; i < n; i++ {
        var pi int
		fmt.Scan(&pi)
		
		if !stop {
			if gap < len(list) {
				for i := pi - gap - 1; i < (pi + gap) && !stop; i++ {
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

	duration := time.Since(start)
	fmt.Fprintln(os.Stderr, duration)
    
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

// package main

// import "fmt"
// import "sort"
// import "math"

// func main() {
//     var N int
//     fmt.Scan(&N)
    
//     var P = make([]int, N)
    
//     for i := 0; i < N; i++ {
//         fmt.Scan(&P[i])
//     }

//     sort.Ints(P);
//     min := float64(P[1] - P[0])
//     for i := 0; i < N - 1; i++ {
//     	min = math.Min(min, float64(P[i + 1] - P[i]));
//     }
//     fmt.Println(min);
// }
