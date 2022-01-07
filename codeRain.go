package main

import (
	"fmt"
	"go-term/ansi"
)

func drawGrid(grid, oldGrid [][]uint8) {
	for y := 0; y < gridSize; y++ {
		for x := 0; x < gridSize; x++ {
			if grid[x][y] != oldGrid[x][y] {
				//offX := (14-x)*3
				//offY := 14-y
				//fmt.Printf("%s%s%3d%s%s", ansi.CUU(offY), ansi.CUB(offX), grid[x][y], ansi.CUF(offX - 3), ansi.CUD(offY))
				//fmt.Printf("%s%s%3d%s", ansi.CUD(y), ansi.CUF(x*3), grid[x][y], ansi.RCP)
				if y > 0 {
					fmt.Print(ansi.CUD(y))
				}
				if x > 0 {
					fmt.Print(ansi.CUF(x*3))
				}
				ansiColor := ansi.RGB(grid[x][y]*5, 0, 0)
				fmt.Printf("%s%3d", ansiColor, grid[x][y])
				fmt.Print(ansi.RCP)
			}
		}
	}
}

func codeRain() {
	
}
