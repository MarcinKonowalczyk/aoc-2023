package day19

import (
	"fmt"
)

func main_2(lines []string) (n int, err error) {
	workflow_lines, part_lines, err := cutAtLine(lines, "")
	if err != nil {
		return 0, err
	}
	var workflows []Workflow = make([]Workflow, len(workflow_lines))
	for i, line := range workflow_lines {
		workflow, err := ParseWorkflowLine(line)
		if err != nil {
			return -1, err
		}
		workflows[i] = workflow
	}

	var parts []Part = make([]Part, len(part_lines))
	for i, line := range part_lines {
		part, err := ParsePartLine(line)
		if err != nil {
			return -1, err
		}
		parts[i] = part
	}

	fmt.Println("Workflows:")
	for _, workflow := range workflows {
		fmt.Println(" ", workflow)
	}

	fmt.Println("Parts:")
	for _, part := range parts {
		fmt.Println(" ", part)
	}

	return 0, nil
}
