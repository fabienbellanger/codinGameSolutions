package main

import (
	"fmt"
	"os"
	"bufio"
	"math"
	"strconv"
)

const Happy = ":)"
const Unhappay = ":("

var unhappyNumbers = [8]int{4, 16, 20, 37, 42, 58, 89, 145}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Buffer(make([]byte, 1000000), 1000000)

    var N int
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&N)
    
    for i := 0; i < N; i++ {
        scanner.Scan()
		x := scanner.Text()
		n, _ := strconv.Atoi(x)
		println(n, " => ", getSum(n, digitsNumber(n)))
    }
    
    // fmt.Fprintln(os.Stderr, "Debug messages...")
    fmt.Println("23 " + Happy)// Write answer to stdout
}

func digitsNumber(i int) int {
	return int(math.Floor(math.Log10(float64(i)) / math.Log10(10.0))) + 1
}

func getSum(i int, n int) int {
	s := 0
	r := i

	for j := n; j > 0; j-- {
		s += (r / int(math.Pow10(j-1))) * (r / int(math.Pow10(j-1)))
		
		r = r % int(math.Pow10(j-1))
	}

	return s
}
