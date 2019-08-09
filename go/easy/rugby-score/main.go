package main

import "fmt"
import "sort"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

var VALUES = [...]int{7, 5, 3}

func main() {
    var N int
	fmt.Scan(&N)

	results := make([][3]int, 0)
	var p = [3]int{}
    
	combinaison(N, p, &results, 0)
	results2 := make([][3]int, len(results), len(results))
	for k, v := range results {
		results2[k][0] = v[0] + v[1]
		results2[k][1] = v[0]
		results2[k][2] = v[2]
	}

	results2 = sortResults(results2)
    
    for _, r := range results2 {
		fmt.Printf("%d %d %d\n", r[0], r[1], r[2])
	}
}

func combinaison(n int, p [3]int, r *[][3]int, i int) {
	if n < 0 {
		return
	}

	if i == len(VALUES) {
		return
	}

	if n == 0 {
		*r = append(*r, p)
		return
	}

	p[i]++
	combinaison(n - VALUES[i], p, r, i)
	p[i]--

	combinaison(n, p, r, i + 1)
}

func sortResults(r [][3]int) [][3]int {
	sort.Slice(r, func(i, j int) bool {
		if r[i][0] < r[j][0] {
			return true
		} else if r[i][0] == r[j][0] {
			if r[i][1] < r[j][1] {
				return true
			} else if r[i][1] == r[j][1] {
				if r[i][2] < r[j][2] {
					return true
				}
			}
		}

		return false
	})

	return r
}
