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
			index := 26
			if letter >= 'A' && letter <= 'Z' {
				index = int(letter - 'A')
			}
			
			for j := i * L; j < i * L + L; j++ {
				fmt.Print(letters[index][j])
			}
		}
		fmt.Println("")
	}
}
