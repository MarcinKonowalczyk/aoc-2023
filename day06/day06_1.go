package day06

import (
	"aoc2023/utils"
	"fmt"
	"math"
	"regexp"
)

func Main(part int, lines []string) (n int, err error) {
	if part == 1 {
		return main_1(lines)
	} else if part == 2 {
		return main_2(lines)
	} else {
		return -1, fmt.Errorf("invalid part")
	}
}

func main_1(lines []string) (n int, err error) {
	times, distances, err := parseLines(lines)
	if err != nil {
		return -1, err
	}
	fmt.Printf("times: %v\n", times)
	fmt.Printf("distances: %v\n", distances)

	prod_N_times := 1
	for i := range times {
		min_time, max_time, err := calcHoldTimeRange(times[i], distances[i])
		if err != nil {
			return -1, err
		}
		N_times := max_time - min_time + 1
		fmt.Printf("min_time: %d, max_time: %d, N_times: %d\n", min_time, max_time, N_times)
		prod_N_times *= N_times
	}

	return prod_N_times, nil
}

var TIMES_LINE_RE = regexp.MustCompile(`Time:(?P<times> *(\d+ *)+)`)

func parseTimesLine(line string) (times []int, err error) {
	matches := utils.GetNamedSubexpsCompiledRe(TIMES_LINE_RE, line)
	return utils.StringOfNumbersToNumbers(matches["times"])
}

var DISTANCES_LINE_RE = regexp.MustCompile(`Distance:(?P<distances> *(\d+ *)+)`)

func parseDistancesLine(line string) (distances []int, err error) {
	matches := utils.GetNamedSubexpsCompiledRe(DISTANCES_LINE_RE, line)
	return utils.StringOfNumbersToNumbers(matches["distances"])
}

func parseLines(lines []string) (times []int, distances []int, err error) {
	if len(lines) != 2 {
		return nil, nil, fmt.Errorf("invalid input")
	}
	times, err = parseTimesLine(lines[0])
	if err != nil {
		return nil, nil, err
	}
	distances, err = parseDistancesLine(lines[1])
	if err != nil {
		return nil, nil, err
	}
	return times, distances, nil
}

// t*n - n^2 = d
// (-d) + (t)*n + (-1)*n^2 = 0
// (-1)*n^2 + (t)*n + (-d) = 0
// det = t**2 - 4*d
// 0.5 * (t +/- sqrt(det))
// t/2 +/- sqrt(det)/2

func calcHoldTimeRange(time int, distance int) (int, int, error) {
	int_determinant := time*time - 4*distance

	if int_determinant <= 0 {
		return -1, -1, fmt.Errorf("invalid determinant")
	}

	var min_time int
	var max_time int

	int_det_sqrt, err := perfectSquareRoot(int_determinant)

	if err == nil {
		// determinant is a perfect square
		min_time = (time-int_det_sqrt)/2 + 1
		max_time = (time+int_det_sqrt)/2 - 1
	} else {
		float_determinant := float64(int_determinant)
		float_time := float64(time)
		float_det_sqrt := math.Sqrt(float_determinant)
		min_time = int((float_time-float_det_sqrt)/2) + 1
		max_time = int((float_time + float_det_sqrt) / 2)
	}

	return min_time, max_time, nil
}

// Naive implementation because it still relies on floating point arithmetic,
// but it should be good enough for this problem.
func perfectSquareRoot(n int) (int, error) {
	sqrt := math.Sqrt(float64(n))
	if sqrt == math.Floor(sqrt) {
		return int(sqrt), nil
	} else {
		return -1, fmt.Errorf("not a perfect square")
	}
}
