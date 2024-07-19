package utils

import "fmt"

type CycleDetection struct {
	Sequence []int // The sequence of elements
	tortoise int   // The position of the tortoise
	hare     int   // The position of the hare
	Start    int   // The start of the cycle
	Period   int   // The period of the cycle
}

func NewCycleDetection() *CycleDetection {
	return &CycleDetection{
		// Sequence: []int{},
		tortoise: -1,
		hare:     -1,
		Start:    -1,
		Period:   -1,
	}
}

func (cd CycleDetection) hasCycle() bool {
	return cd.Start != -1
}

func (cd CycleDetection) String() string {
	return fmt.Sprintf("Tortoise: %d, Hare: %d, Start: %d, Period: %d", cd.tortoise, cd.hare, cd.Start, cd.Period)
}

// Add a new element to the sequence and recomputes the tortoise and hare
func (cd *CycleDetection) Feed(n int) {
	cd.Sequence = append(cd.Sequence, n)

	// Hare always moves
	cd.hare += 1

	if cd.hasCycle() {
		// If we have a cycle saved
		// Check that the cycle which we have found is still valid based on the

		expected_index := cd.Start + (cd.hare-cd.Start)%cd.Period
		expected_value := cd.Sequence[expected_index]
		if n != expected_value {
			// The cycle we're on is invalid.
			cd.Start = -1
			cd.Period = -1
		}
	}

	// Tortoise moves if the sequence is odd
	if len(cd.Sequence)%2 == 1 {
		cd.tortoise += 1
	}

	if cd.Sequence[cd.tortoise] == cd.Sequence[cd.hare] {
		start, period := processTortoiseHare(cd.tortoise, cd.hare, cd.Sequence)
		cd.Start = start
		cd.Period = period
	}

}

func (cd *CycleDetection) FeedAll(seq []int) {
	for _, n := range seq {
		cd.Feed(n)
	}
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

func DetectCycles(seq []int) (int, int) {
	cd := NewCycleDetection()
	cd.FeedAll(seq)
	return cd.Start, cd.Period
}

func (cd *CycleDetection) Reset() {
	cd.Sequence = []int{}
	cd.tortoise = -1
	cd.hare = -1
	cd.Start = -1
	cd.Period = -1
}

func (cd *CycleDetection) Extrapolate(n int) int {
	// Given the current cycle, extrapolate the value at position n
	if !cd.hasCycle() {
		panic("No cycle detected")
	}
	return ExtrapolateCycle(cd.Sequence, n, cd.Start, cd.Period)
}

func ExtrapolateCycle(seq []int, n, start, period int) int {
	if n < start {
		return seq[n]
	}
	delta := n - start
	index := start + delta%period
	return seq[index]
}
