package main

import (
    "fmt"
    "strconv"
    "strings"
)

func main() {
    var N int
    fmt.Scan(&N)

    results := make(map[int]int)
    lines := make([][3]string, N, N)

    for i := 0; i < N; i++ {
        var operation, arg1, arg2 string
        fmt.Scan(&operation, &arg1, &arg2)
        lines[i] = [3]string{operation, arg1, arg2}
    }

    n := 0
    for ; !isMapFull(results, N) && n < 2*N; {
        for i := 0; i < N; i++ {
            operation := lines[i][0]
            arg1 := lines[i][1]
            arg2 := lines[i][2]

            if operation == "VALUE" {
                if _, present := results[i]; !present {
                    var present1 bool
                    var val1 int

                    // Valeur ou variable ?
                    // --------------------
                    if strings.IndexRune(arg1, '$') == -1 {
                        val1, _ = strconv.Atoi(arg1)
                        present1 = true
                    } else {
                        arg1i, _ := strconv.Atoi(arg1[1:])
                        val1, present1 = results[arg1i]
                    }

                    if present1 {
                        results[i] = val1
                    }
                }
            } else {
                if _, present := results[i]; !present {
                    var present1, present2 bool
                    var val1, val2 int

                    // Valeur ou variable ?
                    // --------------------
                    if strings.IndexRune(arg1, '$') == -1 {
                        val1, _ = strconv.Atoi(arg1)
                        present1 = true
                    } else {
                        arg1i, _ := strconv.Atoi(arg1[1:])
                        val1, present1 = results[arg1i]
                    }
                    if strings.IndexRune(arg2, '$') == -1 {
                        val2, _ = strconv.Atoi(arg2)
                        present2 = true
                    } else {
                        arg2i, _ := strconv.Atoi(arg2[1:])
                        val2, present2 = results[arg2i]
                    }

                    if present1 && present2 {
                        if operation == "ADD" {
                            results[i] = val1 + val2
                        } else if operation == "SUB" {
                            results[i] = val1 - val2
                        } else {
                            results[i] = val1 * val2
                        }
                    }
                }
            }
        }

        n++
    }

    for i := 0; i < N; i++ {
        fmt.Println(results[i])
    }
}

func isMapFull(results map[int]int, size int) bool {
    return len(results) == size
}
