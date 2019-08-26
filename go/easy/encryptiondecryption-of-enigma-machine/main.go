package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	scanner.Scan()
	operation := scanner.Text()

	var pseudoRandomNumber uint8
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &pseudoRandomNumber)

	rotors := make([]string, 4)
	rotors[0] = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for i := 0; i < 3; i++ {
		scanner.Scan()
		rotors[i+1] = scanner.Text()
	}
	scanner.Scan()
	message := scanner.Text()

	fmt.Fprintln(os.Stderr, operation, pseudoRandomNumber, rotors, message)

	result := ""
	messageNumber := len(message)
	indexes := make([]int, messageNumber)

	if operation == "ENCODE" {
		// Encode
		// ------
		result = ""
		for j := 0; j < messageNumber; j++ {
			result += string((message[j]-65+pseudoRandomNumber+uint8(j))%26 + 65)
		}

		for i := 0; i < 3; i++ {
			for j := 0; j < messageNumber; j++ {
				indexes[j] = strings.IndexByte(rotors[0], result[j])
			}

			result = ""
			for j := 0; j < messageNumber; j++ {
				result += string(rotors[i+1][indexes[j]])
			}
		}
	} else {
		// Decode
		// ------
		result1 := message
		for i := 3; i >= 0; i-- {
			for j := 0; j < messageNumber; j++ {
				indexes[j] = strings.IndexByte(rotors[i], result1[j])
			}

			result1 = ""
			for j := 0; j < messageNumber; j++ {
				result1 += string(rotors[0][indexes[j]])
			}
		}

		result = ""
		for j := 0; j < messageNumber; j++ {
			result += string((result1[j]+65-pseudoRandomNumber-uint8(j))%26 + 65)
		}
	}

	println(96%26 + 65)

	fmt.Println(result)
}
