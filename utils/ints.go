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

func AbsDiffAndSignTernary[T Integer](a, b T) (T, int) {
	diff := AbsDiff(a, b)
	sign := SignTernary2(a, b)
	return diff, sign
}

func AbsDiffAndSignBinary[T Integer](a, b T) (T, int) {
	diff := AbsDiff(a, b)
	sign := SignBinary2(a, b)
	return diff, sign
}

func IntMax[T Integer](a T, b T) T {
	if a > b {
		return a
	}
	return b
}

func IntMin[T Integer](a T, b T) T {
	if a < b {
		return a
	}
	return b
}

func IntClamp[T Integer](a T, min T, max T) T {
	if a < min {
		return min
	}
	if a > max {
		return max
	}
	return a
}

func BoolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func BoolToSign(b bool) int {
	if b {
		return 1
	}
	return -1
}
