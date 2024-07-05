package day12

import "fmt"

func main_2(lines []string) (n int, err error) {
	parsed_lines := make([]Line, len(lines))
	for line_index, line := range lines {
		parsed_line, err := parseLine(line)
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
		fmt.Println(unfolded_lines[line_index])
	}

	return -1, nil

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
	new_l, err := parseLine(unfolded + " " + new_groups_string)
	if err != nil {
		return Line{}, err
	}
	new_l.orignal_springs = l.orignal_springs
	return new_l, nil
}
