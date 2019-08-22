package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Buffer(make([]byte, 1000000), 1000000)

    // n: the number of temperatures to analyse
    var n int
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&n)
    
    scanner.Scan()
    inputs := strings.Split(scanner.Text()," ")

    var zero int64 = 100000
    for i := 0; i < n; i++ {
        // t: a temperature expressed as an integer ranging from -273 to 5526
        t,_ := strconv.ParseInt(inputs[i],10,32)

        if t == 0 {
            zero = t
            break
        } else if t > 0 {
            if t <= zero || t == -zero {
                zero = t
            }
        } else {
            if zero > 0 {
                if zero + t > 0 {
                    zero = t
                }
            } else {
                if t > zero {
                    zero = t
                }
            }
        }
    }

    if n == 0 {
        zero = 0
    }
    
    // fmt.Fprintln(os.Stderr, "Debug messages...")
    fmt.Printf("%v\n", zero)
}
