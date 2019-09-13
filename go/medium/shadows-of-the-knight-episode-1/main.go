package main

import (
	"fmt"
	"os"
)

type grid struct {
	X1 int
	Y1 int
	X2 int
	Y2 int
}

func main() {
	// W: width of the building.
	// H: height of the building.
	var W, H int
	fmt.Scan(&W, &H)

	// N: maximum number of turns before game over.
	var N int
	fmt.Scan(&N)

	var X0, Y0 int
	fmt.Scan(&X0, &Y0)

	// Rectangle de recherche
	grid := grid{0, 0, W - 1, H - 1}
	fmt.Fprintf(os.Stderr, "grid: %v)\n", grid)

	i := N
	x := X0
	y := Y0
	for {
		// bombDir: the direction of the bombs from batman's current location (U, UR, R, DR, D, DL, L or UL)
		var bombDir string
		fmt.Scan(&bombDir)

		if bombDir == "U" {
			grid.Y2 = y - 1
			y = (y - 1 + grid.Y1) / 2
		} else if bombDir == "D" {
			grid.Y1 = y + 1
			y = (y + 1 + grid.Y2) / 2
		} else if bombDir == "R" {
			grid.X1 = x + 1
			x = (x + 1 + grid.X2) / 2
		} else if bombDir == "L" {
			grid.X2 = x - 1
			x = (x - 1 + grid.X1) / 2
		} else if bombDir == "UR" {
			grid.X1 = x + 1
			grid.Y2 = y - 1
			x = (x + 1 + grid.X2) / 2
			y = (y - 1 + grid.Y1) / 2
		} else if bombDir == "UL" {
			grid.X2 = x - 1
			grid.Y2 = y - 1
			x = (x - 1 + grid.X1) / 2
			y = (y - 1 + grid.Y1) / 2
		} else if bombDir == "DR" {
			grid.X1 = x + 1
			grid.Y1 = y + 1
			x = (x + 1 + grid.X2) / 2
			y = (y + 1 + grid.Y2) / 2
		} else if bombDir == "DL" {
			grid.X2 = x - 1
			grid.Y1 = y + 1
			x = (x - 1 + grid.X1) / 2
			y = (y + 1 + grid.Y2) / 2
		}

		fmt.Fprintf(os.Stderr, "grid: %v)\n", grid)

		// the location of the next window Batman should jump to.
		fmt.Printf("%d %d\n", x, y)

		// Arret de la boucle
		i--
		if i == 0 {
			break
		}
	}
}
