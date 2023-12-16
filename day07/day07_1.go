package day07

import (
	"aoc2023/utils"
	"fmt"
	"sort"
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
	hands := make([]Hand, 0)
	for _, line := range lines {
		hand, err := parseLine(line)
		if err != nil {
			return 0, err
		}
		hands = append(hands, hand)
	}
	orderHands(hands)
	total_score := 0
	for i, hand := range hands {
		rank := i + 1
		score := hand.bid * rank
		total_score += score
	}
	return total_score, nil
}

type handType int

type Hand struct {
	cards     string
	hand_type handType
	bid       int
}

const (
	HIGH_CARD handType = iota
	ONE_PAIR
	TWO_PAIR
	THREE_OF_A_KIND
	FULL_HOUSE
	FOUR_OF_A_KIND
	FIVE_OF_A_KIND
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

func orderHands(hands []Hand) {
	sort.Slice(hands, func(i, j int) bool {
		hand_1 := hands[i]
		hand_2 := hands[j]
		if hand_1.hand_type != hand_2.hand_type {
			return hand_1.hand_type < hand_2.hand_type
		} else {
			for i := 0; i < 5; i++ {
				if hand_1.cards[i] != hand_2.cards[i] {
					ith_card_1 := rune(hand_1.cards[i])
					ith_card_2 := rune(hand_2.cards[i])
					return cardComparison(ith_card_1, ith_card_2)
				}
			}
			return false
		}
	})
}

var card_order = []rune{
	'2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A',
}

func cardComparison(card_1, card_2 rune) bool {
	i1 := utils.ArrayIndexOf(card_order, card_1)
	i2 := utils.ArrayIndexOf(card_order, card_2)
	if i1 == -1 || i2 == -1 {
		panic("invalid card")
	}
	return i1 < i2
}
