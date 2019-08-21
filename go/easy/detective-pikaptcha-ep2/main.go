package main

import (
    "fmt"
    "os"
)

func main() {
    var width, height int
    fmt.Scan(&width, &height)
    
    grid := make([][]byte, 0)
    
    startX := 0
    startY := 0
    for i := 0; i < height; i++ {
        var line string
		fmt.Scan(&line)
		grid = append(grid, []byte(line))
    }

    // Position de départ
    // ------------------
    startFound := false
    for i := 0; i < height && !startFound; i++ {
        for j := 0; j < width && !startFound; j++ {
            c := grid[i][j]

            if c == '>' || c == '<' || c == 'v' || c == '^' {
                startX = j
                startY = i
                startFound = true
            }
        }
    }
    
    // Coté
    // ----
    var side string
    fmt.Scan(&side)

    fmt.Fprintf(os.Stderr, "        Start X: %d\n", startX)
    fmt.Fprintf(os.Stderr, "        Start Y: %d\n", startY)
    fmt.Fprintf(os.Stderr, "Start direction: %s\n", string(grid[startX][startY]))
    fmt.Fprintf(os.Stderr, "      Wall side: %s\n\n", side)

    // Priorités des directions
    // ------------------------
    directionsPriorities := map[string]map[byte][]byte{}
    directionsPriorities["L"] = map[byte][]byte{}
    directionsPriorities["L"]['^'] = []byte("<^>v")
    directionsPriorities["L"]['v'] = []byte(">v<^")
    directionsPriorities["L"]['>'] = []byte("^>v<")
    directionsPriorities["L"]['<'] = []byte("v<^>")
    directionsPriorities["R"] = map[byte][]byte{}
    directionsPriorities["R"]['^'] = []byte(">^<v")
    directionsPriorities["R"]['v'] = []byte("<v>^")
    directionsPriorities["R"]['>'] = []byte("v>^<")
    directionsPriorities["R"]['<'] = []byte("^<v>")
    
    n := 0
    stop := false
    x := startX
    y := startY
    d := grid[startX][startY]
    for ; !stop; {
        for _, v := range directionsPriorities[side][d] {
            if isMovePossible(x, y, width, height, v, grid) {
                if v == '>' {
                    x++
                } else if v == '<' {
                    x--
                } else if v == '^' {
                    y--
                } else {
                    y++
                }

                d = v

                // TODO: Incrémenter le compteur de passage

                break
            } else {
                // TODO ?
            }
        }

        // Fin
        // ---
        if x == startX && y == startY {
            stop = true
        }

        // TODO: A supprimer
        if n > 200 {
            break
        }

        n++

        fmt.Fprintf(os.Stderr, " => (%d, %d) | d: %s\n", x, y, string(d))
    }
    fmt.Fprintf(os.Stderr, "\n => Trouvé en %d\n\n", n)
}

func isMovePossible(x, y, w, h int, d byte, grid [][]byte) bool {
    ok := false

    if d == '^' {
        if y > 0 && grid[y-1][x] != '#' {
            ok = true
        }
    } else if d == 'v' {
        if y + 1 < h && grid[y+1][x] != '#' {
            ok = true
        }
    } else if d == '>' {
        if x + 1 < w && grid[y][x+1] != '#' {
            ok = true
        }
    } else {
        if x > 0 && grid[y][x-1] != '#' {
            ok = true
        }
    }

    return ok
}
