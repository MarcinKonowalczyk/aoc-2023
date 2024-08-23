package utils

func Ternary[T any](cond bool, a T, b T) T {
	if cond {
		return a
	}
	return b
}
