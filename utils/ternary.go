package utils

func Ternary[T any](cond bool, a T, b T) T {
	if cond {
		return a
	}
	return b
}

func TernaryFunc[T any](cond bool, a func() T, b func() T) T {
	if cond {
		return a()
	}
	return b()
}
