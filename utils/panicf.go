package utils

import "fmt"

func Panicf(format string, a ...interface{}) {
	file, line := getParentInfo()
	msg := fmt.Sprintf(format, a...)
	panic(fmt.Sprintf("%s:%d: %s", file, line, msg))
}
