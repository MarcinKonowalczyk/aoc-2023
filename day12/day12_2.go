package day12

import "fmt"

func main_2(lines []string) (n int, err error) {
	parsed_lines := make([]Line, len(lines))
	for line_index, line := range lines {
		parsed_line, err := parseInputLine(line)
		if err != nil {
			return -1, err
		}
		parsed_lines[line_index] = parsed_line
	}

	unfolded_lines := make([]Line, len(parsed_lines))
	for line_index, parsed_line := range parsed_lines {
		new_line, err := unfoldLine(parsed_line, 5)
		if err != nil {
			return -1, err
		}
		unfolded_lines[line_index] = new_line
		// fmt.Println(unfolded_lines[line_index])
	}

	sum_counts := 0

	for _, unfolded_line := range unfolded_lines {
		c := recursiveStepFromLeft(unfolded_line, 0)
		fmt.Println(unfolded_line, "->", c)
		sum_counts += c
	}

	return sum_counts, nil

}

func unfoldLine(l Line, times int) (Line, error) {
	unfolded := ""
	for i := 0; i < times; i++ {
		unfolded += l.orignal_springs
		if i < times-1 {
			unfolded += "?"
		}
	}
	new_groups_string := ""
	for i := 0; i < times; i++ {
		for _, group := range l.groups {
			new_groups_string += fmt.Sprintf("%d,", group)
		}
	}
	new_groups_string = new_groups_string[:len(new_groups_string)-1]
	new_groups, err := parseGroups(new_groups_string)
	if err != nil {
		return Line{}, err
	}
	new_springs, err := parseSprings(unfolded)
	if err != nil {
		return Line{}, err
	}
	return Line{
		l.orignal_springs,
		splitSprings(new_springs, OPERATIONAL),
		new_groups,
	}, nil
}
