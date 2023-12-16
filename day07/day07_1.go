package day07

import (
	"aoc2023/utils"
	"fmt"
	"strconv"
	"strings"
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
	fmt.Println("Hello from main_1")
	fmt.Printf("Got %d lines\n", len(lines))
	for _, line := range lines {
		hand, err := parseLine(line)
		if err != nil {
			return 0, err
		}
		fmt.Println(hand)
	}
	return 0, nil
}

type handType int

type Hand struct {
	cards     string
	hand_type handType
	bid       int
}

const (
	FIVE_OF_A_KIND handType = iota
	FOUR_OF_A_KIND
	FULL_HOUSE
	THREE_OF_A_KIND
	TWO_PAIR
	ONE_PAIR
	HIGH_CARD
)

func parseHandType(cards string) handType {
	if len(cards) != 5 {
		panic("invalid hand")
	}
	count_map := handCount(cards)

	if len(count_map) == 1 {
		return FIVE_OF_A_KIND
	}
	values := utils.MapValues(count_map)

	if len(count_map) == 2 {
		if utils.ArrayIndexOf(values, 4) != -1 {
			// 4 and 1
			return FOUR_OF_A_KIND
		} else {
			// 3 and 2
			return FULL_HOUSE
		}
	} else if len(count_map) == 3 {
		if utils.ArrayIndexOf(values, 3) != -1 {
			// Thee somewhere in the values
			return THREE_OF_A_KIND
		} else {
			// Five cards and three counts. Must be two pairs.
			return TWO_PAIR
		}
	} else if len(count_map) == 4 {
		// Two cards with three counts. Must be one pair.
		return ONE_PAIR
	} else {
		// All cards are unique
		return HIGH_CARD
	}
}

func handCount(cards string) map[rune]int {
	count := make(map[rune]int)
	for _, card := range cards {
		count[card]++
	}
	return count
}

func parseLine(line string) (hand Hand, err error) {
	parts := strings.Split(line, " ")
	if len(parts) != 2 {
		return Hand{}, fmt.Errorf("invalid input")
	}

	cards_string := parts[0]
	if len(cards_string) != 5 {
		return Hand{}, fmt.Errorf("invalid hand")
	}

	bid_string := parts[1]
	bid, err := strconv.Atoi(bid_string)
	if err != nil {
		return Hand{}, fmt.Errorf("invalid bid")
	}

	hand_type := parseHandType(cards_string)

	return Hand{
		cards_string,
		hand_type,
		bid,
	}, nil
}
