package utils

import (
	"testing"
)

func TestCycleDetection_TortoiseHareMoves(t *testing.T) {
	cd := NewCycleDetection()

	// Initial state
	AssertEqual(t, cd.tortoise, -1)
	AssertEqual(t, cd.hare, -1)

	// Both move to 0 initially
	cd.Feed(1)
	AssertEqual(t, cd.tortoise, 0)
	AssertEqual(t, cd.hare, 0)

	// Then tortoise stays still, hare moves to 1
	cd.Feed(1)
	AssertEqual(t, cd.tortoise, 0)
	AssertEqual(t, cd.hare, 1)

	// And now tortoise moves again
	cd.Feed(1)
	AssertEqual(t, cd.tortoise, 1)
	AssertEqual(t, cd.hare, 2)

	// Etc, etc
	cd.Feed(1)
	AssertEqual(t, cd.tortoise, 1)
	AssertEqual(t, cd.hare, 3)
}

func TestCycleDetection_NoCycle(t *testing.T) {
	seq := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	cd := NewCycleDetection()
	for _, n := range seq {
		cd.Feed(n)
	}

	// Because of how the algorithm works, we always have a cycle of 0, 0
	AssertEqualWithComparator(t, cd.cycle, []int{-1, -1}, CompareArrays)
}

func TestCycleDetection_SimpleCycle(t *testing.T) {
	seq := []int{1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3}
	// Cycle 1         ^........^
	// Cycle 2                  ^.................^

	cd := NewCycleDetection()
	for _, n := range seq {
		cd.Feed(n)
	}

	AssertEqualWithComparator(t, cd.cycle, []int{0, 3}, CompareArrays)
}

func TestCycleDetection_CycleWithPrefix(t *testing.T) {
	seq := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3}
	// Cycle                                       ^...................................^

	cd := NewCycleDetection()
	for _, n := range seq {
		cd.Feed(n)
	}

	AssertEqualWithComparator(t, cd.cycle, []int{10, 3}, CompareArrays)
}

func TestCycleDetection_GetResult(t *testing.T) {
	seq := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3}
	// Cycle                                       ^...................................^
	// Start at 10, period is 3

	cd := NewCycleDetection()
	cd.FeedAll(seq)

	start, period := cd.GetResult()
	AssertEqual(t, start, 10)
	AssertEqual(t, period, 3)
}

func TestCycleDetection_GetResult_PrimePeriod(t *testing.T) {
	seq := []int{9, 9, 9, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
	// Cycle              ^...........^

	cd := NewCycleDetection()
	cd.FeedAll(seq)

	start, period := cd.GetResult()
	AssertEqual(t, start, 3)
	AssertEqual(t, period, 5)
}

func TestCycleDetection_GetResult_Temp(t *testing.T) {
	cd := NewCycleDetection()
	start, period := cd.GetResult()
	// Nothing happened yet
	AssertEqual(t, start, -1)
	AssertEqual(t, period, -1)

	cd.Feed(9)
	// 9 -> 0,1
	start, period = cd.GetResult()
	// We've fed one number so far, so the best guess is that it's a cycle of 0, 1
	AssertEqual(t, start, 0)
	AssertEqual(t, period, 1)

	cd.Feed(9)
	// 99 -> 0,1
	start, period = cd.GetResult()
	// This shoudl start looking like a cycle
	AssertEqual(t, start, 0)
	AssertEqual(t, period, 1)

	cd.Feed(9)
	// 999 -> 0,1
	start, period = cd.GetResult()
	// Still a cycle
	AssertEqual(t, start, 0)
	AssertEqual(t, period, 1)

	cd.Feed(1)
	// 9991 -> 3,1
	start, period = cd.GetResult()
	// Ups. I guess not. This invalidates the cycle
	AssertEqual(t, start, -1)
	AssertEqual(t, period, -1)
}
