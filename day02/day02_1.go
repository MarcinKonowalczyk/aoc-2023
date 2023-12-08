package day02

import (
	"errors"
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
	sum_game_ids := 0
	for line_index, line := range lines {
		line, err = simplifyLine(line)
		if err != nil {
			return -1, err
		}
		games, err := lineToGames(line)
		if err != nil {
			return -1, err
		}
		possible := areGamesPossible(games)
		// fmt.Println(games, possible)
		if possible {
			game_id := line_index + 1
			sum_game_ids += game_id
		}
	}
	return sum_game_ids, nil
}

type game struct {
	red   int
	green int
	blue  int
}

func lineToGames(line string) ([]game, error) {
	game_strings := strings.Split(line, ";")
	var out []game = make([]game, len(game_strings))
	for game_index, game_string := range game_strings {
		marble_strings := strings.Split(game_string, ",")
		for _, marble_string := range marble_strings {
			marble_string = strings.Trim(marble_string, " ")
			parts := strings.Split(marble_string, " ")
			if len(parts) != 2 {
				return out, errors.New("len(parts) != 2)")
			}
			n, err := strconv.Atoi(parts[0])
			if err != nil {
				return out, err
			}
			switch parts[1] {
			case "red":
				out[game_index].red = n
			case "green":
				out[game_index].green = n
			case "blue":
				out[game_index].blue = n
			default:
				return out, fmt.Errorf("invalid marble specifier %s", parts[1])
			}
		}
	}
	return out, nil
}

func simplifyLine(line string) (string, error) {
	if !strings.HasPrefix(line, "Game ") {
		return "", errors.New("invalid line format - the line does not start with 'Game '")
	}
	parts := strings.Split(line, ":")
	if len(parts) != 2 {
		return "", errors.New("invalid line format - the line has more than one ':'")
	}
	line = parts[1]
	line = strings.Trim(line, " ")
	// line = strings.ReplaceAll(line, "red", "r")
	// line = strings.ReplaceAll(line, "green", "g")
	// line = strings.ReplaceAll(line, "blue", "b")
	return line, nil
}

const N_RED = 12
const N_GREEN = 13
const N_BLUE = 14

func areGamesPossible(games []game) bool {
	for _, game := range games {
		if game.red > N_RED {
			return false
		}
		if game.green > N_GREEN {
			return false
		}
		if game.blue > N_BLUE {
			return false
		}
	}
	return true
}
