package utils

func AbsDiff[T Integer](a T, b T) T {
	if a > b {
		return a - b
	}
	return b - a
}

// Return the absolute value of a number.
func Abs[T Integer](a T) T {
	if a < 0 {
		return -a
	}
	return a
}

// Return the sign of a number as -1, 0, or 1.
func SignTernary1[T Integer](a T) int {
	if a < 0 {
		return -1
	}
	if a > 0 {
		return 1
	}
	return 0
}

func SignTernary2[T Integer](a, b T) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}

// Return the sign of a number as -1 or 1.
// Zero is considered positive.
func SignBinary1[T Integer](a T) int {
	if a < 0 {
		return -1
	}
	return 1
}

func SignBinary2[T Integer](a, b T) int {
	if a < b {
		return -1
	}
	return 1
}
