package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Buffer(make([]byte, 1000000), 1000000)

    var L int
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&L)
    
    var H int
    scanner.Scan()
	fmt.Sscan(scanner.Text(),&H)
    
    scanner.Scan()
	text := strings.ToUpper(scanner.Text())
	fmt.Fprintln(os.Stderr, "Text:", text)

	// Alphabet en ASCII Art
	// ---------------------
	var letters [27][]string
	lettersMatch := map[string]int{
		"A": 0,
		"B": 1,
		"C": 2,
		"D": 3,
		"E": 4,
		"F": 5,
		"G": 6,
		"H": 7,
		"I": 8,
		"J": 9,
		"K": 10,
		"L": 11,
		"M": 12,
		"N": 13,
		"O": 14,
		"P": 15,
		"Q": 16,
		"R": 17,
		"S": 18,
		"T": 19,
		"U": 20,
		"V": 21,
		"W": 22,
		"X": 23,
		"Y": 24,
		"Z": 25,
	}
	
    for i := 0; i < H; i++ {
        scanner.Scan()
		row := scanner.Text()

		k := -1
		l := 0
		for j := 0; j < 27 * L; j++ {
			if j % L == 0 {
				k++
				l = 0
			}

			letters[k] = append(letters[k], string(row[j]))

			l++
		}		
	}

	// Affichage du texte en ASCII Art
	// -------------------------------
	for i := 0; i < H; i++ {
		for _, letter := range text {
			letterStr := string(letter)

			index := 26
			if i, ok := lettersMatch[letterStr]; ok {
				index = i
			}
			
			for j := i * L; j < i * L + L; j++ {
				fmt.Print(letters[index][j])
			}
		}
		fmt.Println("")
	}
}
