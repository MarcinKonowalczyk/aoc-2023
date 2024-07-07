package day12

import (
	"fmt"
	"strings"
)

func main_2(lines []string) (n int, err error) {

	unfolded_lines := make([]string, len(lines))
	for line_index, line := range lines {
		new_line, err := unfoldLine(line, 5)
		if err != nil {
			return -1, err
		}
		unfolded_lines[line_index] = new_line
		// fmt.Println(unfolded_lines[line_index])
	}

	parsed_lines := make([]Line, len(lines))
	for line_index, line := range unfolded_lines {
		parsed_line, err := parseInputLine(line)
		if err != nil {
			return -1, err
		}
		parsed_lines[line_index] = parsed_line
	}

	sum_counts := 0

	for _, parsed_line := range parsed_lines {
		c := recursiveStepFromLeft(parsed_line, 0)
		fmt.Println(parsed_line, "->", c)
		sum_counts += c
	}

	return sum_counts, nil

}

func unfoldLine(line string, times int) (string, error) {
	parts := strings.Split(line, " ")
	if len(parts) != 2 {
		return "", fmt.Errorf("invalid line")
	}
	springs_string := parts[0]
	groups_string := parts[1]

	unfolded := ""
	for i := 0; i < times; i++ {
		unfolded += springs_string
		if i < times-1 {
			unfolded += "?"
		}
	}
	unfolded += " "
	for i := 0; i < times; i++ {
		unfolded += groups_string
		if i < times-1 {
			unfolded += ","
		}
	}
	return unfolded, nil
}
