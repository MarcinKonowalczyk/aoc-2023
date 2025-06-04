package cycledetector

import (
	"testing"

	"github.com/MarcinKonowalczyk/assert"
)

func TestCycleDetection_TortoiseHareMoves(t *testing.T) {
	cd := NewCycleDetection()

	// Initial state
	assert.Equal(t, cd.tortoise, -1)
	assert.Equal(t, cd.hare, -1)

	// Both move to 0 initially
	cd.Feed(1)
	assert.Equal(t, cd.tortoise, 0)
	assert.Equal(t, cd.hare, 0)

	// Then tortoise stays still, hare moves to 1
	cd.Feed(1)
	assert.Equal(t, cd.tortoise, 0)
	assert.Equal(t, cd.hare, 1)

	// And now tortoise moves again
	cd.Feed(1)
	assert.Equal(t, cd.tortoise, 1)
	assert.Equal(t, cd.hare, 2)

	// Etc, etc
	cd.Feed(1)
	assert.Equal(t, cd.tortoise, 1)
	assert.Equal(t, cd.hare, 3)
}

func TestCycleDetection_NoCycle(t *testing.T) {
	seq := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	cd := NewCycleDetection()
	for _, n := range seq {
		cd.Feed(n)
	}

	// Because of how the algorithm works, we always have a cycle of 0, 0
	// assert.Equal(ithComparator(t, cd.Cycle, []int{-1, -1}, CompareArrays)
	assert.Equal(t, cd.hasCycle(), false)
}

func TestCycleDetection_SimpleCycle(t *testing.T) {
	seq := []int{1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3}
	// Cycle 1         ^........^
	// Cycle 2                  ^.................^

	cd := NewCycleDetection()
	for _, n := range seq {
		cd.Feed(n)
	}
	assert.Equal(t, cd.hasCycle(), true)
	assert.Equal(t, cd.Start, 0)
	assert.Equal(t, cd.Period, 3)
}

func TestCycleDetection_CycleWithPrefix(t *testing.T) {
	seq := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3}
	// Cycle                                       ^...................................^

	cd := NewCycleDetection()
	for _, n := range seq {
		cd.Feed(n)
	}

	// assert.Equal(ithComparator(t, cd.Cycle, []int{10, 3}, CompareArrays)
	assert.Equal(t, cd.hasCycle(), true)
	assert.Equal(t, cd.Start, 10)
	assert.Equal(t, cd.Period, 3)
}

func TestCycleDetection_PrimePeriod(t *testing.T) {
	seq := []int{9, 9, 9, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
	// Cycle              ^...........^

	cd := NewCycleDetection()
	cd.FeedAll(seq)

	assert.Equal(t, cd.Start, 3)
	assert.Equal(t, cd.Period, 5)
}

func TestCycleDetection_SpecificCase(t *testing.T) {
	cd := NewCycleDetection()
	// Nothing happened yet
	assert.Equal(t, cd.Start, -1)
	assert.Equal(t, cd.Period, -1)

	cd.Feed(9)
	// 9 -> 0,1
	// We've fed one number so far, so the best guess is that it's a cycle of 0, 1
	assert.Equal(t, cd.Start, 0)
	assert.Equal(t, cd.Period, 1)

	cd.Feed(9)
	// 99 -> 0,1
	// This shoudl start looking like a cycle
	assert.Equal(t, cd.Start, 0)
	assert.Equal(t, cd.Period, 1)

	cd.Feed(9)
	// 999 -> 0,1
	// Still a cycle
	assert.Equal(t, cd.Start, 0)
	assert.Equal(t, cd.Period, 1)

	cd.Feed(1)
	// 9991 -> 3,1
	// Ups. I guess not. This invalidates the cycle
	assert.Equal(t, cd.Start, -1)
	assert.Equal(t, cd.Period, -1)
}

func TestCycleDetection_Extrapolate(t *testing.T) {
	seq := []int{9, 9, 9, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
	// Cycle              ^...........^

	cd := NewCycleDetection()
	cd.FeedAll(seq)

	for i := cd.Start; i < len(seq); i++ {
		assert.Equal(t, seq[i], cd.Extrapolate(i))
	}
}

func TestDetectCycle(t *testing.T) {
	seq := []int{4, 4, 5, 4, 4, 4, 5, 4, 4, 4, 5, 4, 4, 4, 5, 4, 4, 4, 5, 4, 4, 4, 5, 4, 4, 4, 5, 4, 4}
	start, period := DetectCycle(seq)
	assert.Equal(t, start, 0)
	assert.Equal(t, period, 4)
}
