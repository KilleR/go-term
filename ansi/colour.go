package ansi

import (
	"fmt"
	"github.com/gookit/color"
)
//ESC[ 38;2;⟨r⟩;⟨g⟩;⟨b⟩ m
func RGB(r,g,b uint8) string {
	clr := color.RgbTo256(r,g,b)

	return fmt.Sprintf("\u001b[38;5;%dm", clr)

	return fmt.Sprintf("\u001b[38;2;%d;%d:%d m", r,g,b)
}

func ColorCode(c int) string {
	return fmt.Sprintf("\u001b[38;5;%dm", c)
}