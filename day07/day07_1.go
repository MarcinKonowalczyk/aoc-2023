package day07

import (
	"aoc2023/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Main(part int, lines []string, verbose bool) (n int, err error) {
	if part == 1 {
		return main_1(lines, verbose)
	} else if part == 2 {
		return main_2(lines, verbose)
	} else {
		return -1, fmt.Errorf("invalid part")
	}
}

func main_1(lines []string, verbose bool) (n int, err error) {
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

type Hand struct {
	cards string
	bid   int
}

type handType float32

const (
	HIGH_CARD handType = iota
	ONE_PAIR
	TWO_PAIR
	THREE_OF_A_KIND
	FULL_HOUSE
	FOUR_OF_A_KIND
	FIVE_OF_A_KIND
)

func handTypeToString(hand_type handType) string {
	switch hand_type {
	case HIGH_CARD:
		return "HIGH_CARD"
	case ONE_PAIR:
		return "ONE_PAIR"
	case TWO_PAIR:
		return "TWO_PAIR"
	case THREE_OF_A_KIND:
		return "THREE_OF_A_KIND"
	case FULL_HOUSE:
		return "FULL_HOUSE"
	case FOUR_OF_A_KIND:
		return "FOUR_OF_A_KIND"
	case FIVE_OF_A_KIND:
		return "FIVE_OF_A_KIND"
	default:
		panic("invalid hand type")
	}
}

func cardsToType(cards string) handType {
	if len(cards) != 5 {
		panic("invalid hand")
	}
	count_map := handCount(cards)

	if len(count_map) == 1 {
		return FIVE_OF_A_KIND
	}
	_, values := utils.MapKeysAndValues(count_map)

	if len(count_map) == 2 {
		if utils.ArrayContains(values, 4) {
			// 4 and 1
			return FOUR_OF_A_KIND
		} else {
			// 3 and 2
			return FULL_HOUSE
		}
	} else if len(count_map) == 3 {
		if utils.ArrayContains(values, 3) {
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

func handCount(cards string) map[string]int {
	count := make(map[string]int)
	for _, card := range cards {
		count[string(card)]++
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

	return Hand{
		cards_string,
		bid,
	}, nil
}

func orderHands(hands []Hand) {
	sort.Slice(hands, func(i, j int) bool {
		hand_1_cards := hands[i].cards
		hand_2_cards := hands[j].cards
		hand_type_1 := cardsToType(hand_1_cards)
		hand_type_2 := cardsToType(hand_2_cards)
		if hand_type_1 != hand_type_2 {
			return hand_type_1 < hand_type_2
		} else {
			for i := 0; i < 5; i++ {
				ith_card_1 := rune(hand_1_cards[i])
				ith_card_2 := rune(hand_2_cards[i])
				if ith_card_1 != ith_card_2 {
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
