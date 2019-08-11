package main

import (
	"fmt"
	// "os"
)

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
    var N int
	fmt.Scan(&N)
	
    /*
    Triangle height = N
    Triangle width  = 2N - 1

	Height = 2N
	Width  = 4N - 1
    Middle = 2N - 1
    
    N = 1
    -----
    .*
    * *

    N = 3
    -----

	.    *
		***
	   *****
	  *     *
	 ***   ***
    ***** *****
    
	*/
	var line string

	for i := 0; i < 2*N; i++ {
		line = ""
        
        for j := 0; j < 2 * N + i; j++ {
            // fmt.Fprintf(os.Stderr, "%d x %d\n", i, j)
            if j < 2 * N - 1 - i {
                line += " "
            } else if i < N {
                // First triangle
                // --------------
                line += "*"
                // if j < ((2 * N - 1) - i) + (2 * i + 1) {
                //     line += "*"
                // } else {
                //     line += "o"
                // }
            } else {
                // Others triangles
                // ----------------
                line += "*"
            }
        }

        if i == 0 {
            line = "." + line[1:]
        }

		fmt.Println(line)
	}
    
    // fmt.Fprintln(os.Stderr, "Debug messages...")
    fmt.Println("answer")// Write answer to stdout
}
