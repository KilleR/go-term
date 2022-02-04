package main

import (
	"fmt"
	"go-term/ansi"
	"math/rand"
	"time"
)

const gridHeight = 15
const gridWidth = 80
const oldAge = 20
const maxAge = 40

type activeRune struct {
	x, y int
}

type gridRune struct {
	rune
	age int
}

func drawRain(grid [][]*gridRune) {
	fmt.Print(ansi.RCP)
	for y := 0; y < terminal.height; y++ {
		for x := 0; x < terminal.width; x++ {
			if grid[x][y] == nil {
				continue
			}
			cellRune := grid[x][y]
			if cellRune.age == 0 || cellRune.age == 1 || cellRune.age > maxAge || cellRune.age > oldAge {
				if y > 0 {
					fmt.Print(ansi.CUD(y))
				}
				if x > 0 {
					fmt.Print(ansi.CUF(x))
				}

				var ansiColor string
				switch cellRune.age {
				case 0:
					ansiColor = ansi.RGB(0, 255, 255)
				case 1:
					ansiColor = ansi.RGB(0, 255, 0)
				default:
					if cellRune.age > oldAge {
						ansiColor = ansi.RGB(0, 150, 0)
					}
				}
				if cellRune.age > maxAge {
					fmt.Printf("%s%s", ansiColor, " ")
					grid[x][y] = nil
				} else {
					fmt.Printf("%s%s", ansiColor, string(cellRune.rune))
				}
				fmt.Print(ansi.RCP)
			}
			cellRune.age++
		}
	}
}

func codeRain() {
	defer fmt.Print(ansi.CUD(terminal.height))

	var activeRunes []*activeRune

	grid := make([][]*gridRune, terminal.width)
	for x := 0; x < terminal.width; x++ {
		grid[x] = make([]*gridRune, terminal.height)
	}
	for i := 0; i < terminal.height; i++ {
		fmt.Println()
	}
	fmt.Print(ansi.CUU(terminal.height), ansi.SCP)

	for {
		//oldGrid := make([][]*gridRune, terminal.width)
		//for x, g := range grid {
		//	oldGrid[x] = make([]*gridRune, terminal.height)
		//	for y, val := range g {
		//		oldGrid[x][y] = val
		////	}
		//}

		for i, ar := range activeRunes {
			ar.y = ar.y + 1
			if ar.y >= terminal.height {
				activeRunes = append(activeRunes[:i], activeRunes[i+1:]...)
				continue
			}
			grid[ar.x][ar.y] = &gridRune{letterRunes[rand.Intn(len(letterRunes))], 0}

		}
		activeRunes = append(activeRunes, &activeRune{
			x: rand.Intn(terminal.width),
			y: 0,
		})
		activeRunes = append(activeRunes, &activeRune{
			x: rand.Intn(terminal.width),
			y: 0,
		})

		drawRain(grid)
		time.Sleep(time.Millisecond * 50)
	}
}
