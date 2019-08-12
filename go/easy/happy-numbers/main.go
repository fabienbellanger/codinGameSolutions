package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

const Happy = ":)"
const Unhappy = ":("

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Buffer(make([]byte, 1000000), 1000000)

    var N int
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&N)
    
    for i := 0; i < N; i++ {
        scanner.Scan()
		x := scanner.Text()
		
		s := x
		for {
			s = getSum(s)
			
			if s == "1" {
				fmt.Printf("%s %s\n", x, Happy)
				break
			} else if s == "4" || s == "16" || s == "20" || s == "37" || s == "42" || s == "58" || s == "89" || s == "145" {
				fmt.Printf("%s %s\n", x, Unhappy)
				break
			}
		}
    }
}

func getSum(s string) string {
	var sum2 int

	for _, value := range s {
		num, _ := strconv.Atoi(string(value))
		
		sum2 += int(num) * int(num)
	}
	
	return strconv.Itoa(sum2)
}
