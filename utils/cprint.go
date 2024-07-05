package utils

import "fmt"

type Color int

// Foreground text colors
const (
	Black Color = iota + 30
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

func (c Color) String() string {
	return fmt.Sprintf("\033[%dm", c)
}

const Reset = "\033[0m"

// func cprint(color Color, format string, a ...interface{}) {
// 	fmt.Printf(color.String()+format+Reset, a...)
// }

func Csprintf(color Color, format string, a ...interface{}) string {
	return color.String() + fmt.Sprintf(format, a...) + Reset
}

func Cprintf(color Color, format string, a ...interface{}) {
	fmt.Printf(Csprintf(color, format, a...))
}

func Cboolf(b bool) string {
	if b {
		return Csprintf(Green, "true")
	} else {
		return Csprintf(Red, "false")
	}
}
