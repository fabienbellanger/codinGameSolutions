package main

import (
    "fmt"
    "os"
    "bufio"
    "regexp"
    "strconv"
)

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Buffer(make([]byte, 1000000), 1000000)

    var n int
    scanner.Scan()
    fmt.Sscan(scanner.Text(), &n)

    invalidISBN := make([][]byte, 0)
    
    for i := 0; i < n; i++ {
        scanner.Scan()
        // isbn := scanner.Text()
        isbn := scanner.Bytes()
        ok := checkISBN(isbn)
        if !ok {
            invalidISBN = append(invalidISBN, isbn)
        }
    }
    
    // Output
    // ------
    fmt.Printf("%d invalid:\n", len(invalidISBN))
    for _, val := range invalidISBN {
        fmt.Printf("%s\n", val)
    }
}

// checkISBN checks if ISBN is correct
func checkISBN(s []byte) bool {
    // Width of string equals to 10 or 13
    if len(s) != 10 && len(s) != 13 {
        return false
    }

    // String only have number
    if match, _ := regexp.Match(`^[\dX]+$`, s); !match {
        return false
    }

    if len(s) == 10 {
        // ISBN 10
        // -------
        sum := 0
        for i := 0; i < 9; i++ {
            num, err := strconv.Atoi(string(s[i]))
            if err != nil {
                return false
            }

            sum += num * (10 - i)
        }

        var controlDigit string
        if sum % 11 == 0 {
            controlDigit = "0"
        } else {
            if 11 - sum % 11 == 10 {
                controlDigit = "X"
            } else {
                controlDigit = strconv.Itoa(11 - sum % 11)
            }
        }
        lastDigit := string(s[9])

        if controlDigit != lastDigit {
            return false
        }
    } else {
        // ISBN 13
        // -------
        sum := 0
        for i := 0; i < 12; i++ {
            num, err := strconv.Atoi(string(s[i]))
            if err != nil {
                return false
            }

            if i % 2 == 0 {
                sum += num
            } else {
                sum += num * 3
            }
        }
        
        controlDigit := 10 - sum % 10
        if controlDigit == 10 {
            controlDigit = 0
        }
        lastDigit, err := strconv.Atoi(string(s[12]))
        
        if err != nil {
            return false
        }

        if controlDigit != lastDigit {
            return false
        }
    }

    return true
}
