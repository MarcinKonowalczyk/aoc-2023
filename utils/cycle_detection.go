package utils

import "fmt"

type CycleDetection struct {
	Sequence   []int // The sequence of elements
	tortoise   int   // The position of the tortoise
	hare       int   // The position of the hare
	cycle      []int // The cycle [start, end]
	prev_cycle []int // The previous cycle [start, end]
}

func NewCycleDetection() *CycleDetection {
	return &CycleDetection{
		// Sequence: []int{},
		tortoise:   -1,
		hare:       -1,
		cycle:      []int{-1, -1},
		prev_cycle: []int{-1, -1},
	}
}

func (cd CycleDetection) hasCycle() bool {
	return cd.cycle[0] != -1
}

func (cd CycleDetection) hasPrevCycle() bool {
	return cd.prev_cycle[0] != -1
}

func (cd CycleDetection) String() string {
	return fmt.Sprintf("Tortoise: %d, Hare: %d, Cycle: %v", cd.tortoise, cd.hare, cd.cycle)
}

// Add a new element to the sequence and recomputes the tortoise and hare
func (cd *CycleDetection) Feed(n int) {
	cd.Sequence = append(cd.Sequence, n)

	// Hare always moves
	cd.hare += 1

	if cd.hasCycle() {
		// If we have a cycle saved
		// Check that the cycle which we have found is still valid based on the

		expected_index := cd.cycle[0] + (cd.hare-cd.cycle[0])%cd.cycle[1]
		expected_value := cd.Sequence[expected_index]
		if n != expected_value {
			// The cycle we're on is invalid.
			cd.cycle = []int{-1, -1}
		}
	}

	// Tortoise moves if the sequence is odd
	if len(cd.Sequence)%2 == 1 {
		cd.tortoise += 1
	}

	if cd.Sequence[cd.tortoise] == cd.Sequence[cd.hare] {
		start, period := processTortoiseHare(cd.tortoise, cd.hare, cd.Sequence)
		cd.cycle = []int{start, period}
		cd.prev_cycle = cd.cycle
	}

}

func (cd *CycleDetection) FeedAll(seq []int) {
	for _, n := range seq {
		cd.Feed(n)
	}
}

func (cd *CycleDetection) GetResult() (int, int) {
	return cd.cycle[0], cd.cycle[1]
}

func processTortoiseHare(
	tort, hare int, seq []int) (start int, period int) {
	if tort == 0 && hare == 0 {
		// This is the start. The best guess for the cycle is the whole sequence
		if len(seq) != 1 {
			panic("tortoise and hare are both 0 but len(seq) != 1")
		}
		return 0, 1
	}

	period = hare - tort
	start = tort

	// But we might have walked multiple cycles! Let's check all the divisors of the period
	// and see if any of them are valid, starting from smallest to largest
	// The first we hit is a valid cycle
	divisors := Divisors(period)
	if len(divisors) <= 2 {
		// Period is prime. This is the shortest cycle
	} else {
		for _, d := range divisors {
			if d == 1 {
				// Ignore 1
				continue
			}
			// Check if it is a valid cycle
			valid := true
			for i := 0; i < d; i++ {
				if seq[start+i] != seq[start+i+d] {
					valid = false
					break
				}
			}
			if valid {
				// Accept this period and since we are iterating from smallest to largest
				// this is the shortest cycle
				period = d
				break
			}
		}
	}

	end := start + period
	// Ok. Now we want ot find the start of the cycle. We might have skipped multiple cycles or part of the cycle
	// Walk start and end backwards until they no longer match
	for {
		if start == 0 {
			break
		}
		if seq[start] == seq[end] {
			start--
			end--
		} else {
			break
		}
	}

	if start == 0 && end == len(seq)-1 {
		// Special case where the cycle is the entire sequence
		// fmt.Println("special_case", "start", start, "period", end-start)
		return start, end - start
	}

	if seq[start] != seq[end] {
		// We have walked too far
		start++
		end++
	}

	if seq[start] != seq[end] {
		// Sanity check
		panic("Something went wrong")
	}

	// fmt.Println("final", "start", start, "period", end-start)

	return start, end - start
}
