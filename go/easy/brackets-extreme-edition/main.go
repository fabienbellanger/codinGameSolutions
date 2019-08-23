package main

import "fmt"

func main() {
	var expression string
	fmt.Scan(&expression)

	chars := ""
	isValid := true

	for _, char := range expression {
		if char == '{' || char == '[' || char == '(' {
			chars += string(char)
		} else {
			if char == '}' || char == ']' || char == ')' {
				validChar(char, &chars, &isValid)
				if !isValid {
					break
				}
			}
		}
	}

	if isValid && len(chars) != 0 {
		isValid = false
	}

	fmt.Printf("%t\n", isValid)
}

func validChar(c rune, s *string, b *bool) {
	l := len(*s)
	var co byte = '('

	if c == '}' {
		co = '{'
	} else if c == ']' {
		co = '['
	}

	if l == 0 || (*s)[l-1] != co {
		*b = false
	} else {
		*s = (*s)[:l-1]
	}
}
