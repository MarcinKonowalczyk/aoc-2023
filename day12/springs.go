package day12

import (
	"fmt"
)

type Spring int

const (
	OPERATIONAL Spring = iota
	DAMAGED
	UNKNOWN
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
	orignal_springs string     // original springs string
	blocks          [][]Spring // blocks of springs split at OPERATIONAL
	groups          []int      // grounps of consecutive DAMAGED springs
}

func (l Line) String() string {
	s := ""
	// s += l.orignal_springs + " ["
	s += "["
	for i, block := range l.blocks {
		s += springsString(block)
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

func (l Line) copy() Line {
	copy_blocks := make([][]Spring, len(l.blocks))
	for i, block := range l.blocks {
		copy_block := make([]Spring, len(block))
		copy(copy_block, block)
		copy_blocks[i] = copy_block
	}
	copy_groups := make([]int, len(l.groups))
	copy(copy_groups, l.groups)
	copy_orignal_springs := l.orignal_springs
	return Line{copy_orignal_springs, copy_blocks, copy_groups}
}

func springsString(ss []Spring) string {
	s := ""
	for _, spring := range ss {
		s += spring.String()
	}
	return s
}
