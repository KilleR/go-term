package ansi

import (
	"fmt"
	"github.com/gookit/color"
)

//ESC[ 38;2;⟨r⟩;⟨g⟩;⟨b⟩ m
func RGB(r, g, b uint8) string {
	clr := color.RgbTo256(r, g, b)

	return fmt.Sprintf("\u001b[38;5;%dm", clr)

	return fmt.Sprintf("\u001b[38;2;%d;%d:%d m", r, g, b)
}

func ColorCode(c int) string {
	return fmt.Sprintf("\u001b[38;5;%dm", c)
}

const BLACK = "\u001b[30m"
const RED = "\u001b[31m"
const GREEN = "\u001b[32m"
const YELLOW = "\u001b[33m"
const BLUE = "\u001b[34m"
const MAGENTA = "\u001b[35m"
const CYAN = "\u001b[36m"
const WHITE = "\u001b[37m"
const RESET = "\u001b[0m"
