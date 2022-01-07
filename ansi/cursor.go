package ansi

import "fmt"

func CUU(n int) string {
	return fmt.Sprintf("\u001b[%dA", n)
}
func CUD(n int) string {
	return fmt.Sprintf("\u001b[%dB", n)
}
func CUF(n int) string {
	return fmt.Sprintf("\u001b[%dC", n)
}
func CUB(n int) string {
	return fmt.Sprintf("\u001b[%dD", n)
}

func CUP(x, y int) string {
	return fmt.Sprintf("\u001b[%d;%dH", x, y)
}

const SCP = "\u001b7"
const RCP = "\u001b8"