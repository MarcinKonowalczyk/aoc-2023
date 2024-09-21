package day20

import (
	"aoc2023/utils"
	"fmt"
	"os"
	"sort"
)

func main_2(lines []string, verbose bool) (n int, err error) {
	modules, names, err := ParseModules(lines, verbose)
	if err != nil {
		return 0, err
	}

	// find rx
	rx, ok := modules["rx"]
	if !ok {
		return 0, fmt.Errorf("Module 'rx' not found")
	}

	if len(rx.Inputs()) != 1 {
		return 0, fmt.Errorf("Module 'rx' should have exactly one input")
	}

	// rx's ancestor
	rxi := rx.Inputs()[0]

	// rx's ancestor's subgraphs
	rxii := modules[rxi].Inputs()

	if verbose {
		fmt.Printf("rx's parent is '%s', and its grandparens are '%v'.\n", rxi, rxii)
	}

	// Make the input queues
	input_queues := make(map[string][]element, 0)
	for name := range modules {
		input_queues[name] = make([]element, 0)
	}

	// Send a bunch of test pulses to the broadcaster
	N_TEST_PULSES := 20000
	lows := make([]int, N_TEST_PULSES)
	highs := make([]int, N_TEST_PULSES)

	hooks := make(map[string]Hook)

	// Make a hook for each of the rx's grandparent's
	values := make(map[string][]bool, 0)
	received_pulses := make(map[string]int, 0)
	current_index := 0

	for _, grandparent := range rxii {
		name := grandparent
		values[name] = make([]bool, 0)
		received_pulses[name] = 0
		hooks[name] = func(sender string, value bool) {
			if len(values[name]) <= current_index {
				values[name] = append(values[name], true)
			}
			values[name][current_index] = values[name][current_index] && value
			received_pulses[name]++
			if !value && verbose {
				fmt.Printf("Hook: '%s' got a low pulse\n", name)
			}
		}
	}

	filename := "dump.txt"

	f, _ := os.Create(filename)
	defer f.Close()

	for i := 0; i < N_TEST_PULSES; i++ {
		current_index = i
		input_queues["broadcaster"] = []element{{"", false}}
		n_low, n_high, err := drive_machine(modules, input_queues, names, hooks, false)
		if err != nil {
			return 0, err
		}
		n_low++ // Count the initial pulse

		// _, v := utils.MapKeysAndValuesSorted(received_pulses)
		// str := fmt.Sprintf("%05d: %d\n", i, v)
		// f.WriteString(str)

		for _, grandparent := range rxii {
			received_pulses[grandparent] = 0
		}

		lows[i] = n_low
		highs[i] = n_high
	}

	// For each of the grandparent, count the number of low pulses
	for grandparent, values := range values {
		highs := utils.ArrayReduce(values, 0, func(acc int, value bool) int {
			if value {
				return acc + 1
			}
			return acc
		})
		lows := len(values) - highs
		if lows < 3 {
			return 0, fmt.Errorf("grandparent '%s' did not receive enough low pulses", grandparent)
		}
	}

	// For each grandparent find the index of the first and the second low pulse
	low_indices := make(map[string][]int, len(rxii))

	for grandparent, values := range values {
		first_low := -1
		second_low := -1
		third_low := -1
		for i, value := range values {
			if !value {
				if first_low == -1 {
					first_low = i
				} else if second_low == -1 {
					second_low = i
				} else if third_low == -1 {
					third_low = i
				}
			}
		}
		low_indices[grandparent] = []int{first_low, second_low, third_low}
	}

	for grandparent, indices := range low_indices {
		diff, err := utils.ArrayDiff(indices, 1)
		if err != nil {
			return 0, err
		}
		if diff[0] != diff[1] {
			return 0, fmt.Errorf("grandparent '%s' did not receive pulses at the same time", grandparent)
		}
	}

	// find when the grandparents get in phase
	starts := make([]int, len(rxii))
	periods := make([]int, len(rxii))
	for i, grandparent := range rxii {
		indices := low_indices[grandparent]
		starts[i] = indices[0] + 1
		periods[i] = indices[1] - indices[0]
	}

	// sort the starts and periods
	sort.Ints(starts)
	sort.Ints(periods)

	if verbose {
		fmt.Println("starts:", starts)
		fmt.Println("periods:", periods)
	}

	// Make sure the starts and period are the same
	// This must be true for the simplified version to work
	for i := 1; i < len(starts); i++ {
		if starts[i] != periods[i] {
			return 0, fmt.Errorf("starts and periods are not the same")
		}
	}

	lcm := utils.LCM(periods[0], periods[1])
	for i := 2; i < len(periods); i++ {
		lcm = utils.LCM(lcm, periods[i])
	}

	return lcm, nil
}

// 2971974843540 too low
// 233338595643977 is just right!!!!!!!

func findInPhase(starts, periods []int) (int, error) {
	if len(starts) != len(periods) {
		return 0, fmt.Errorf("starts and periods should have the same length")
	}

	N := len(starts)

	if N == 0 {
		return 0, fmt.Errorf("starts and periods should not be empty")
	}

	if N == 1 {
		// Only one sequence. It is always in phase with itself
		return 0, nil
	}

	q2, x2, err := findInPhase2(starts[2], periods[2], starts[3], periods[3])
	if err != nil {
		return 0, err
	}

	q, x, err := findInPhase2(starts[0], periods[0], starts[1], periods[1])
	if err != nil {
		return 0, err
	}

	if N == 2 {
		return x, err
	}

	fmt.Println("q:", q, "x:", x)
	fmt.Println("q2:", q2, "x2:", x2)

	// We have more than these two sequences. That's ok. We've just found the index when they are in phase.

	new_starts := []int{q}
	new_periods := []int{x}
	for i := 2; i < N; i++ {
		new_starts = append(new_starts, starts[i])
		new_periods = append(new_periods, periods[i])
	}

	fmt.Println("new_starts:", new_starts)
	fmt.Println("new_periods:", new_periods)

	return findInPhase(new_starts, new_periods)
}

func findInPhase2(s1, p1, s2, p2 int) (int, int, error) {

	// https://math.stackexchange.com/questions/2218763/how-to-find-lcm-of-two-numbers-when-one-starts-with-an-offset

	// Make sure s1 < s2
	if s1 > s2 {
		s1, s2 = s2, s1
		p1, p2 = p2, p1
	}

	// Shift back such that 1 starts at 0
	shift1 := s1
	s2 = s2 - shift1
	s1 = s1 - shift1

	fmt.Println("s1:", s1, "p1:", p1, "s2:", s2, "p2:", p2)

	// //  Shift such that 2 starts during the first period of 1
	// shift2 := (s2 / p1) * p1
	// s2 = s2 - shift2

	// fmt.Println("shift1:", shift1, "shift2:", shift2)
	// fmt.Println("s1:", s1, "p1:", p1, "s2:", s2, "p2:", p2)

	// lcm := utils.LCM(periods[0], periods[1])
	// fmt.Println("LCM:", lcm)

	g, s, t := utils.ExtendedGCD(p1, p2)
	fmt.Println("g:", g, "s:", s, "t:", t)

	d := s2 - s1

	fmt.Println("d:", d)
	// Check if d is divisible by g
	if d%g != 0 {
		return -1, -1, fmt.Errorf("d (%d) is not divisible by g (%d)", d, g)
	}

	z := d / g
	// fmt.Println("z:", z)
	m := z * s
	n := -z * t
	if m < 0 {
		m += p2
	}
	if n < 0 {
		n += p1
	}
	// fmt.Println("m:", m, "n:", n)

	chk := (s1 + m*p1) - (s2 + n*p2)
	// fmt.Println("chk:", chk)

	if chk != 0 {
		return -1, -1, fmt.Errorf("check failed: %d", chk)
	}

	lcm := utils.LCM(p1, p2)
	// fmt.Println("lcm:", lcm)

	// Some large x
	x := (m*p1 + s1)
	// Earliest x
	x = x % lcm
	// fmt.Println("x:", x)

	for x < 0 {
		x += lcm
	}
	for x < s1 {
		x += lcm
	}
	for x < s2 {
		x += lcm
	}

	chk_m := (x - s1) % (p1)
	if chk_m != 0 {
		return -1, -1, fmt.Errorf("chk_m failed: %d", chk_m)
	}
	chk_n := (x - s2) % (p2)
	if chk_n != 0 {
		return -1, -1, fmt.Errorf("chk_n failed: %d", chk_n)
	}

	// Combined start
	q := s1
	if q < s2 {
		q = s2
	}

	return q, x, nil

}
