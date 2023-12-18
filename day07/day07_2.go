package day07

import (
	"aoc2023/utils"
	"sort"
)

func main_2(lines []string) (n int, err error) {
	hands := make([]Hand, 0)
	for _, line := range lines {
		hand, err := parseLine(line)
		if err != nil {
			return 0, err
		}
		hands = append(hands, hand)
	}

	orderHandsWithJokers(hands)

	// for _, hand := range hands {
	// 	cards := hand.cards
	// 	hand_type := cardsToTypeWithJokers(cards)
	// 	fmt.Println("cards", cards, "becomes", handTypeToString(hand_type))
	// }

	total_score := 0
	for i, hand := range hands {
		// fmt.Println(hand)
		rank := i + 1
		score := hand.bid * rank
		total_score += score
	}
	return total_score, nil
}

func sortFuncWithJokers(hand_1_cards, hand_2_cards string) bool {
	hand_type_1 := cardsToTypeWithJokers(hand_1_cards)
	hand_type_2 := cardsToTypeWithJokers(hand_2_cards)
	if hand_type_1 != hand_type_2 {
		return hand_type_1 < hand_type_2
	} else {
		// hands have the same type
		for k := 0; k < 5; k++ {
			kth_card_1 := rune(hand_1_cards[k])
			kth_card_2 := rune(hand_2_cards[k])
			if kth_card_1 != kth_card_2 {
				return cardComparisonWithJokers(kth_card_1, kth_card_2)
			}
		}
		panic("duplicate hand")
	}
}

func orderHandsWithJokers(hands []Hand) {
	cmp := func(i, j int) bool {
		hand_i_cards := hands[i].cards
		hand_j_cards := hands[j].cards
		return sortFuncWithJokers(hand_i_cards, hand_j_cards)
	}
	sort.Slice(hands, cmp)
}

func cardsToTypeWithJokers(cards string) handType {
	if len(cards) != 5 {
		panic("invalid hand")
	}
	count_map := handCount(cards)

	keys, values := utils.MapKeysAndValues(count_map)

	index_of_joker := utils.ArrayIndexOf(keys, "J")
	if index_of_joker == -1 {
		// No jokers. Proceed as normal.
		return cardsToType(cards)
	}

	if len(count_map) <= 2 {
		// Two cases:
		// 1) 5 jokers.
		// 2) Some jokers and some number of other cards all of the same type.
		// Either way, we can make it a five of a kind.
		return FIVE_OF_A_KIND
	}

	N_jokers := values[index_of_joker]

	if len(count_map) == 3 {
		if N_jokers >= 2 {
			// 1) 3 jokers and 2 other cards.
			// 2) 2 jokers and 3 other cards.
			return FOUR_OF_A_KIND
		} else { // N_jokers == 1
			if N_jokers != 1 {
				panic("invalid hand")
			}
			// 1 joker and 4 other cards.
			if utils.ArrayContains(values, 3) {
				// the other cards are a three of a kind. We can make four of a kind.
				return FOUR_OF_A_KIND
			} else {
				// the other cards are two pairs. We can make a full house.
				return FULL_HOUSE
			}
		}
	} else if len(count_map) == 4 {
		// 1) 2 jokers and 3 other cards across 3 card types. Best is three of a kind.
		// 2) 1 joker and 4 other cards across 3 card types. There is already a pair so best
		//    is three of a kind.
		return THREE_OF_A_KIND
	} else {
		// All five cards are different, but one is a joker. We can make a pair.
		return ONE_PAIR
	}

}

var card_order_with_jokers = []rune{
	'J', '2', '3', '4', '5', '6', '7', '8', '9', 'T', 'Q', 'K', 'A',
}

func cardComparisonWithJokers(card_1, card_2 rune) bool {
	i1 := utils.ArrayIndexOf(card_order_with_jokers, card_1)
	i2 := utils.ArrayIndexOf(card_order_with_jokers, card_2)
	if i1 == -1 || i2 == -1 {
		panic("invalid card")
	}
	if i1 == i2 {
		panic("same cards")
	}
	return i1 <= i2
}
