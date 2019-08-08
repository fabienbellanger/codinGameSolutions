package main

import "fmt"
import "os"
import "bufio"
import "path/filepath"
import "strings"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Buffer(make([]byte, 1000000), 1000000)

    // N: Number of elements which make up the association table.
    var N int
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&N)
    
    // Q: Number Q of file names to be analyzed.
    var Q int
    scanner.Scan()
	fmt.Sscan(scanner.Text(),&Q)
	
	const UNKNOWN = "UNKNOWN"
	extensionList := make(map[string]string, N)
	result := make([]string, Q, Q)
    
    for i := 0; i < N; i++ {
        var EXT, MT string
        scanner.Scan()
		fmt.Sscan(scanner.Text(),&EXT, &MT)
		extensionList[strings.ToLower(EXT)] = MT
	}

    for i := 0; i < Q; i++ {
        scanner.Scan()
		fileExt := strings.ToLower(filepath.Ext(scanner.Text()))

		// Recherche extention
		// -------------------
		if len(fileExt) == 0 {
			result[i] = UNKNOWN
		} else {
			if v, ok := extensionList[fileExt[1:]]; ok {
				result[i] = v
			} else {
				result[i] = UNKNOWN
			}
		}
	}

	for _, v := range result {
		fmt.Println(v)
	}
}
