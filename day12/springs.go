package day12

import (
	"fmt"
)

type Spring uint8

const (
	UNKNOWN Spring = iota
	DAMAGED
	OPERATIONAL
)

func (s Spring) String() string {
	switch s {
	case OPERATIONAL:
		return "."
	case DAMAGED:
		return "#"
	case UNKNOWN:
		return "?"
	default:
		panic("invalid spring (1)")
	}
}

func fromString(s string) (Spring, error) {
	switch s {
	case ".":
		return OPERATIONAL, nil
	case "#":
		return DAMAGED, nil
	case "?":
		return UNKNOWN, nil
	default:
		return UNKNOWN, fmt.Errorf("invalid spring (2)")
	}
}

type Line struct {
	blocks [][]Spring // blocks of springs split at OPERATIONAL
	groups []uint8    // grounps of consecutive DAMAGED springs
}

func (l Line) String() string {
	s := ""
	s += "["
	for i, block := range l.blocks {
		s += SpringsToString(block)
		if i < len(l.blocks)-1 {
			s += ","
		}
	}
	s += "] "
	for i, group := range l.groups {
		s += fmt.Sprintf("%d", group)
		if i < len(l.groups)-1 {
			s += ","
		}
	}
	return s
}

func (l Line) Hash() uint64 {
	var buf uint64

	// The sate must be non-zero so we start by setting the buffer to 1
	buf = 1

	// For each spring in each block, write its value to the buffer one by one
	// Each spring is either UNKNOWN = 0 or DAMAGED = 1
	// Pack them into the bits of a uint64

	i := 0
	for _, block := range l.blocks {
		for _, spring := range block {
			buf = buf<<2 | uint64(spring)
			i += 2
			if i == 32 {
				// Every 32 bits, we need to multiply the value by a prime and xor the
				// right bits with the left bits
				buf *= 1111111111111111111
				buf ^= buf >> 32
				i = 0
			}
		}
		buf = buf<<2 | 2 // 2 is the separator between blocks
		i += 2
		if i == 32 {
			// Every 32 bits, we need to multiply the value by a prime and xor the
			// right bits with the left bits
			buf *= 1111111111111111111
			buf ^= buf >> 32
			i = 0
		}
	}

	// Push a padding pattern of 10101010 to the buffer to separate the blocks
	buf = buf<<2 | 2
	i += 2
	if i == 32 {
		buf *= 1111111111111111111
		buf ^= buf >> 32
		i = 0
	}

	for _, group := range l.groups {
		// Group is a uint8. Pack its bits into the buffer one by one
		for j := 0; j < 8; j++ {
			buf = buf<<1 | uint64(group>>j&1)
			i++
			if i == 32 {
				buf *= 1111111111111111111
				buf ^= buf >> 32
				i = 0
			}
		}
	}

	if i != 0 {
		buf *= 1111111111111111111
		buf ^= buf >> 32
	}

	return buf
}

func (l Line) copy() Line {
	copy_blocks := make([][]Spring, len(l.blocks))
	for i, block := range l.blocks {
		copy_block := make([]Spring, len(block))
		copy(copy_block, block)
		copy_blocks[i] = copy_block
	}
	copy_groups := make([]uint8, len(l.groups))
	copy(copy_groups, l.groups)
	return Line{copy_blocks, copy_groups}
}

func SpringsToString(ss []Spring) string {
	s := ""
	for _, spring := range ss {
		s += spring.String()
	}
	return s
}

func SpringsFromString(s string) ([]Spring, error) {
	springs := make([]Spring, len(s))
	for i, c := range s {
		spring, err := fromString(string(c))
		if err != nil {
			return nil, fmt.Errorf("invalid spring: %s", string(c))
		}
		springs[i] = spring
	}
	return springs, nil
}

func MustSpringsFromString(s string) []Spring {
	springs, err := SpringsFromString(s)
	if err != nil {
		panic(err)
	}
	return springs
}
