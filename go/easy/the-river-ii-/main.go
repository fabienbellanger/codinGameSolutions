package main

import "fmt"

func main() {
	var r1 int
	fmt.Scan(&r1)

	ok := false
	used := make(map[int]bool)
	for i := 1; i < r1; i++ {
		if _, present := used[i]; !present {
			used[i] = false
			n := i
			for n < r1 {
				n += getDigitSum(n)
				used[n] = false
			}

			if n == r1 {
				ok = true
				break
			}
		}
	}

	if ok {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func getDigitSum(n int) (s int) {
	for n != 0 {
		s = s + n%10
		n = n / 10
	}

	return
}
