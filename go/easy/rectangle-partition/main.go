package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	var w, h, countX, countY int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &w, &h, &countX, &countY)

	scanner.Scan()
	inputs := strings.Split(scanner.Text(), " ")
	for i := 0; i < countX; i++ {
		x, _ := strconv.ParseInt(inputs[i], 10, 32)
		_ = x
	}
	scanner.Scan()
	inputs = strings.Split(scanner.Text(), " ")
	for i := 0; i < countY; i++ {
		y, _ := strconv.ParseInt(inputs[i], 10, 32)
		_ = y
	}

	// fmt.Fprintln(os.Stderr, "Debug messages...")
	fmt.Println("0") // Write answer to stdout
}
