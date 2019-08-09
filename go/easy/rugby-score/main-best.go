package main

import (
	"fmt"
)

func main() {
    var N int
	fmt.Scan(&N)
	
    for t := 0; t <= N/5; t++ {
        for k := 0; k <= t; k++ {
            var R = N - t * 5 - k * 2
            if R >= 0 && R % 3 == 0 {
                fmt.Println( t, k, R / 3 )
            }
        }
    }
}
