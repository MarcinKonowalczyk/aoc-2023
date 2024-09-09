package day15

import (
	"aoc2023/utils"
	"fmt"
	"strconv"
	"strings"
)

type lens struct {
	label string
	value int
}

func (l lens) Hash() uint8 {
	return hash(l.label)
}

func main_2(lines []string, verbose bool) (n int, err error) {
	if len(lines) != 1 {
		return -1, fmt.Errorf("invalid input")
	}
	steps := strings.Split(lines[0], ",")

	boxes := [256][]lens{}
	for _, step := range steps {
		// Check if '=' is in the step
		op, nl, err := parseStep(step)
		if err != nil {
			return -1, err
		}
		box := boxes[nl.Hash()]
		I := utils.ArrayIndexOfFunc(box, func(ol lens) bool { return ol.label == nl.label })
		if op {
			// op is add
			if I == -1 {
				box = append(box, nl)
			} else {
				(box)[I] = nl
			}
		} else {
			// op is delete
			if I != -1 {
				box = utils.ArrayRemoveIndex(box, I)
			}
		}
		boxes[nl.Hash()] = box
	}

	focusing_power := 0
	for i, box := range boxes {
		for j, l := range box {
			focusing_power += (i + 1) * (j + 1) * l.value
		}
	}

	return focusing_power, nil
}

func parseStep(s string) (bool, lens, error) {
	var label string
	var value int
	var op bool
	var err error

	parts := strings.Split(s, "=")
	op = true
	if len(parts) > 2 {
		return false, lens{}, fmt.Errorf("invalid step")
	} else if len(parts) == 2 {
		label = parts[0]
		value, err = strconv.Atoi(parts[1])
		if err != nil {
			return false, lens{}, err
		}
	} else {
		parts = strings.Split(s, "-")
		op = false
		if len(parts) != 2 {
			return false, lens{}, fmt.Errorf("invalid step")
		}
		label = parts[0]
		if parts[1] == "" {
			value = -1
		} else {
			value, err = strconv.Atoi(parts[1])
			if err != nil {
				return false, lens{}, err
			}
		}
	}
	return op, lens{label: label, value: value}, nil
}
