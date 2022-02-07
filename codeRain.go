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
const rippleWidth = 5

var drawX, drawY int

type activeRune struct {
	x, y int
}

type gridRune struct {
	rune
	age int
}

func drawRain(grid [][]*gridRune) {
	for y := 0; y < terminal.height; y++ {
		for x := 0; x < terminal.width; x++ {
			if grid[x][y] == nil {
				continue
			}
			cellRune := grid[x][y]
			if cellRune.age == 0 || cellRune.age == 1 || cellRune.age == maxAge || cellRune.age > oldAge {
				offsetX := x - drawX
				offsetY := y - drawY
				if offsetY > 0 {
					fmt.Print(ansi.CUD(offsetY))
					drawY += offsetY
				}
				if offsetY < 0 {
					fmt.Print(ansi.CUU(-offsetY))
					drawY += offsetY
				}
				if offsetX > 0 {
					fmt.Print(ansi.CUF(offsetX))
					drawX += offsetX
				}
				if offsetX < 0 {
					fmt.Print(ansi.CUB(-offsetX))
					drawX += offsetX
				}

				var ansiColor string
				switch cellRune.age {
				case 0:
					ansiColor = ansi.RGB(0, 255, 255)
				case 1:
					ansiColor = ansi.RGB(0, 255, 0)
				default:
					ansiColor = ansi.RGB(0, 150, 0)
				}

				if cellRune.age > maxAge {
					fmt.Printf("%s%s%s", ansiColor, " ", ansi.CUB(1))
					grid[x][y] = nil
				} else {
					fmt.Printf("%s%s%s", ansiColor, string(cellRune.rune), ansi.CUB(1))
				}
			}
			cellRune.age++
		}
	}
}

func codeRain() {
	defer fmt.Print(ansi.CUD(terminal.height - 1))

	var activeRunes []*activeRune

	grid := make([][]*gridRune, terminal.width)
	for x := 0; x < terminal.width; x++ {
		grid[x] = make([]*gridRune, terminal.height)
	}
	for i := 0; i < terminal.height; i++ {
		fmt.Println()
	}
	fmt.Print(ansi.CUU(terminal.height), ansi.SCP)

	activeRunes = append(activeRunes, &activeRune{
		x: rand.Intn(terminal.width),
		y: 0,
	})

drawLoop:
	for {

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

		drawRain(grid)
		select {
		case <-sigs:
			break drawLoop
		case <-time.After(time.Millisecond * 50):
		}

	}

	//return
	// ripple
	for i := 0; i < 110; i++ {
		rippleMinX := (50 + i/2 - rippleWidth) * terminal.width / 100
		rippleMaxX := (50 + i/2 + rippleWidth) * terminal.width / 100
		rippleMinY := (50 + i/2 - rippleWidth) * terminal.height / 100
		rippleMaxY := (50 + i/2 + rippleWidth) * terminal.height / 100
		ripple2MinX := (50 - i/2 - rippleWidth) * terminal.width / 100
		ripple2MaxX := (50 - i/2 + rippleWidth) * terminal.width / 100
		ripple2MinY := (50 - i/2 - rippleWidth) * terminal.height / 100
		ripple2MaxY := (50 - i/2 + rippleWidth) * terminal.height / 100

		for y := 0; y < terminal.height; y++ {
			for x := 0; x < terminal.width; x++ {
				if grid[x][y] == nil {
					continue
				}
				cellRune := grid[x][y]

				if y > 0 {
					fmt.Print(ansi.CUD(y))
				}
				if x > 0 {
					fmt.Print(ansi.CUF(x))
				}

				var ansiColor string
				inRipple := (rippleMinX < x && x < rippleMaxX && y < rippleMaxY && ripple2MinY < y) ||
					(rippleMinY < y && y < rippleMaxY && x < rippleMaxX && ripple2MinX < x) ||
					(ripple2MinX < x && x < ripple2MaxX && y < rippleMaxY && ripple2MinY < y) ||
					(ripple2MinY < y && y < ripple2MaxY && x < rippleMaxX && ripple2MinX < x)
				switch age := cellRune.age; {
				case age == 1 && inRipple:
					ansiColor = ansi.WHITE
				case age == 1, age > 1 && age <= oldAge && inRipple:
					ansiColor = ansi.RGB(0, 255, 255)
				case age > 1 && age <= oldAge, age > oldAge && inRipple:
					ansiColor = ansi.RGB(0, 255, 0)
				case age > oldAge:
					ansiColor = ansi.RGB(0, 150, 0)
				}

				fmt.Printf("%s%s", ansiColor, string(cellRune.rune))
				fmt.Print(ansi.RCP)
			}
		}

	}
	time.Sleep(time.Millisecond * 10)

	boxTop := terminal.height/2 - 2
	boxLeft := terminal.width/2 - 9
	fmt.Printf("%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s", ansi.CUD(boxTop), ansi.CUF(boxLeft), "╔════════════════╗", ansi.CUB(18), ansi.CUD(1), "║                ║", ansi.CUB(18), ansi.CUD(1), "║  SYSTEM ERROR  ║", ansi.CUB(18), ansi.CUD(1), "║                ║", ansi.CUB(18), ansi.CUD(1), "╚════════════════╝")

}
