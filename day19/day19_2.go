package day19

import (
	"fmt"
)

func main_2(lines []string, verbose bool) (n int, err error) {
	workflow_lines, _, err := cutAtLine(lines, "")
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

	if verbose {
		fmt.Println("Workflows:")
		for _, workflow := range workflows {
			fmt.Println(" ", workflow)
		}
	}

	return 0, nil
}
