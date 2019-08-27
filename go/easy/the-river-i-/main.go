package main

import "fmt"

func main() {
	var r1 int
	fmt.Scan(&r1)

	var r2 int
	fmt.Scan(&r2)

	for r1 != r2 {
		if r1 > r2 {
			r2 += getDigitSum(r2)
		} else if r2 > r1 {
			r1 += getDigitSum(r1)
		} else {
			break
		}
	}

	fmt.Println(r1)
}

func getDigitSum(n int) (s int) {
	for n != 0 {
		s = s + n%10
		n = n / 10
	}

	return
}
