package main

import (
	"fmt"
	"go-term/ansi"
	"math/rand"
	"time"
)

const gridHeight = 15
const gridWidth = 80

type activeRune struct {
	x, y int
}

func drawRain(grid, oldGrid [][]rune) {
	fmt.Print(ansi.RCP)
	for y := 0; y < gridHeight; y++ {
		for x := 0; x < gridWidth; x++ {
			if grid[x][y] != oldGrid[x][y] {
				//offX := (14-x)*3
				//offY := 14-y
				//fmt.Printf("%s%s%3d%s%s", ansi.CUU(offY), ansi.CUB(offX), grid[x][y], ansi.CUF(offX - 3), ansi.CUD(offY))
				//fmt.Printf("%s%s%3d%s", ansi.CUD(y), ansi.CUF(x*3), grid[x][y], ansi.RCP)
				if y > 0 {
					fmt.Print(ansi.CUD(y))
				}
				if x > 0 {
					fmt.Print(ansi.CUF(x))
				}
				ansiColor := ansi.RGB(0, 255, 0)
				fmt.Printf("%s%s", ansiColor, string(grid[x][y]))
				fmt.Print(ansi.RCP)
			}
		}
	}
}

func codeRain() {
	defer fmt.Print(ansi.CUD(gridHeight))

	var activeRunes []*activeRune

	grid := make([][]rune, gridWidth)
	for x := 0; x < gridWidth; x++ {
		grid[x] = make([]rune, gridHeight)
	}
	for i := 0; i < gridHeight; i++ {
		fmt.Println()
	}
	fmt.Print(ansi.CUU(gridHeight), ansi.SCP)

	for {
		oldGrid := make([][]rune, gridWidth)
		for x, g := range grid {
			oldGrid[x] = make([]rune, gridHeight)
			for y, val := range g {
				oldGrid[x][y] = val
			}
		}

		for i, ar := range activeRunes {
			ar.y = ar.y + 1
			if ar.y >= gridHeight {
				activeRunes = append(activeRunes[:i], activeRunes[i+1:]...)
				continue
			}
			grid[ar.x][ar.y] = letterRunes[rand.Intn(len(letterRunes))]

		}
		activeRunes = append(activeRunes, &activeRune{
			x: rand.Intn(gridWidth),
			y: 0,
		})

		drawRain(grid, oldGrid)
		time.Sleep(time.Millisecond * 100)
	}
}
