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

    scanner.Scan()
    message := scanner.Text()
    res := binary(message)

    series := make([]string, 0)
    currentChar := ""
    currentPos := -1

    for _, char := range res {
        if string(char) != currentChar {
            series = append(series, string(char))
            currentChar = string(char)
            currentPos++
        } else {
            series[currentPos] += string(char)
        }
    }

    for i, v := range series {
        if i == 0 {
            if v == "0" {
                fmt.Printf("00 0")
            } else {
                fmt.Printf("0 0")
            }
        }  else {
            if strings.Contains(v, "0") {
                fmt.Printf(" 00 ")
            } else {
                fmt.Printf(" 0 ")
            }

            fmt.Printf(strings.Repeat("0", len(v)))
        }
    }
    fmt.Printf("\n")
}

func binary(s string) string {
    res := ""
    for _, c := range s {
        res = fmt.Sprintf("%s%.7b", res, c)
    }
    return res
}
