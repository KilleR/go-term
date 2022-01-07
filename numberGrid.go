package main

import (
	"fmt"
	"go-term/ansi"
	"math/rand"
	"time"
)

const gridSize = 15

func printGrid(grid [][]uint8) {
	for y := 0; y < gridSize; y++ {
		for x := 0; x < gridSize; x++ {
			fmt.Printf("%3d", grid[x][y])
		}
		if y < gridSize-1 {
			fmt.Print("\n")
		}
	}
	fmt.Printf("%s%s%s", ansi.CUU(gridSize-1), ansi.CUB(gridSize*3), ansi.SCP)
}

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

func numberGrid() {
	var grid [][]uint8
	var oldGrid [][]uint8

	grid = make([][]uint8, gridSize)
	for i, _ := range grid {
		grid[i] = make([]uint8, gridSize)
	}
	oldGrid = make([][]uint8, gridSize)
	for i, _ := range oldGrid {
		oldGrid[i] = make([]uint8, gridSize)
	}

	printGrid(grid)

	for n:= 0; n < 10000; n++ {
		time.Sleep(1 * time.Millisecond)
		x := rand.Intn(gridSize)
		y := rand.Intn(gridSize)

		grid[x][y]++
		drawGrid(grid, oldGrid)
		for i, _ := range grid {
			copy(oldGrid[i], grid[i])
		}
	}
	fmt.Printf("%s%s\n", ansi.CUD(gridSize), ansi.CUF(gridSize*3))
}
