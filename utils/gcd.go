package utils

// https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int) int {

	result := a * b / GCD(a, b)

	// for i := 0; i < len(integers); i++ {
	// 	result = LCM(result, integers[i])
	// }

	return result
}

// in addition to GCD(a, b), computes the coefficients of Bézout's identity,
// which are integers x and y such that ax + by = gcd(a, b)
// https://en.wikipedia.org/wiki/Extended_Euclidean_algorithm
// Returns gcd(a, b), x, y
func ExtendedGCD(a, b int) (int, int, int) {
	x := 1
	s := 0
	y := 0
	t := 1

	for b != 0 {
		quotient := a / b
		a, b = b, a-quotient*b
		x, s = s, x-quotient*s
		y, t = t, y-quotient*t
	}

	return a, x, y
}
