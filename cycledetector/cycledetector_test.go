package cycledetector

import (
	"aoc2023/utils"
	"testing"
)

func TestCycleDetection_TortoiseHareMoves(t *testing.T) {
	cd := NewCycleDetection()

	// Initial state
	utils.AssertEqual(t, cd.tortoise, -1)
	utils.AssertEqual(t, cd.hare, -1)

	// Both move to 0 initially
	cd.Feed(1)
	utils.AssertEqual(t, cd.tortoise, 0)
	utils.AssertEqual(t, cd.hare, 0)

	// Then tortoise stays still, hare moves to 1
	cd.Feed(1)
	utils.AssertEqual(t, cd.tortoise, 0)
	utils.AssertEqual(t, cd.hare, 1)

	// And now tortoise moves again
	cd.Feed(1)
	utils.AssertEqual(t, cd.tortoise, 1)
	utils.AssertEqual(t, cd.hare, 2)

	// Etc, etc
	cd.Feed(1)
	utils.AssertEqual(t, cd.tortoise, 1)
	utils.AssertEqual(t, cd.hare, 3)
}

func TestCycleDetection_NoCycle(t *testing.T) {
	seq := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	cd := NewCycleDetection()
	for _, n := range seq {
		cd.Feed(n)
	}

	// Because of how the algorithm works, we always have a cycle of 0, 0
	// utils.AssertEqualWithComparator(t, cd.Cycle, []int{-1, -1}, CompareArrays)
	utils.AssertEqual(t, cd.hasCycle(), false)
}

func TestCycleDetection_SimpleCycle(t *testing.T) {
	seq := []int{1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3}
	// Cycle 1         ^........^
	// Cycle 2                  ^.................^

	cd := NewCycleDetection()
	for _, n := range seq {
		cd.Feed(n)
	}
	utils.AssertEqual(t, cd.hasCycle(), true)
	utils.AssertEqual(t, cd.Start, 0)
	utils.AssertEqual(t, cd.Period, 3)
}

func TestCycleDetection_CycleWithPrefix(t *testing.T) {
	seq := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3}
	// Cycle                                       ^...................................^

	cd := NewCycleDetection()
	for _, n := range seq {
		cd.Feed(n)
	}

	// utils.AssertEqualWithComparator(t, cd.Cycle, []int{10, 3}, CompareArrays)
	utils.AssertEqual(t, cd.hasCycle(), true)
	utils.AssertEqual(t, cd.Start, 10)
	utils.AssertEqual(t, cd.Period, 3)
}

func TestCycleDetection_PrimePeriod(t *testing.T) {
	seq := []int{9, 9, 9, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
	// Cycle              ^...........^

	cd := NewCycleDetection()
	cd.FeedAll(seq)

	utils.AssertEqual(t, cd.Start, 3)
	utils.AssertEqual(t, cd.Period, 5)
}

func TestCycleDetection_SpecificCase(t *testing.T) {
	cd := NewCycleDetection()
	// Nothing happened yet
	utils.AssertEqual(t, cd.Start, -1)
	utils.AssertEqual(t, cd.Period, -1)

	cd.Feed(9)
	// 9 -> 0,1
	// We've fed one number so far, so the best guess is that it's a cycle of 0, 1
	utils.AssertEqual(t, cd.Start, 0)
	utils.AssertEqual(t, cd.Period, 1)

	cd.Feed(9)
	// 99 -> 0,1
	// This shoudl start looking like a cycle
	utils.AssertEqual(t, cd.Start, 0)
	utils.AssertEqual(t, cd.Period, 1)

	cd.Feed(9)
	// 999 -> 0,1
	// Still a cycle
	utils.AssertEqual(t, cd.Start, 0)
	utils.AssertEqual(t, cd.Period, 1)

	cd.Feed(1)
	// 9991 -> 3,1
	// Ups. I guess not. This invalidates the cycle
	utils.AssertEqual(t, cd.Start, -1)
	utils.AssertEqual(t, cd.Period, -1)
}

func TestCycleDetection_Extrapolate(t *testing.T) {
	seq := []int{9, 9, 9, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
	// Cycle              ^...........^

	cd := NewCycleDetection()
	cd.FeedAll(seq)

	for i := cd.Start; i < len(seq); i++ {
		utils.AssertEqual(t, seq[i], cd.Extrapolate(i))
	}
}

func TestDetectCycle(t *testing.T) {
	seq := []int{4, 4, 5, 4, 4, 4, 5, 4, 4, 4, 5, 4, 4, 4, 5, 4, 4, 4, 5, 4, 4, 4, 5, 4, 4, 4, 5, 4, 4}
	start, period := DetectCycle(seq)
	utils.AssertEqual(t, start, 0)
	utils.AssertEqual(t, period, 4)
}
