package day19

import (
	"fmt"
	"time"
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

func cutAtLine(lines []string, header string) ([]string, []string, error) {
	header_index := -1
	for i := 0; i < len(lines); i++ {
		if lines[i] == header {
			header_index = i
			break
		}
	}
	if header_index == -1 {
		return lines, nil, fmt.Errorf("header not found: %s", header)
	}
	return lines[:header_index], lines[header_index+1:], nil
}

var CH_MAP map[string]chan Part
var REJECTED chan Part
var ACCEPTED chan Part

const MAX_PARTS = 1

func main_1(lines []string, verbose bool) (n int, err error) {
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

	if verbose {
		fmt.Println("Workflows:")
		for _, workflow := range workflows {
			fmt.Println(" ", workflow)
		}

		fmt.Println("Parts:")
		for _, part := range parts {
			fmt.Println(" ", part)
		}
	}

	// Create part input channels
	CH_MAP = make(map[string]chan Part)
	for _, workflow := range workflows {
		CH_MAP[workflow.name] = make(chan Part, MAX_PARTS)
	}

	REJECTED = make(chan Part, MAX_PARTS)
	ACCEPTED = make(chan Part, MAX_PARTS)

	// Start workflow servers
	for _, workflow := range workflows {
		ch := CH_MAP[workflow.name]
		go workflow_server(workflow, ch, ACCEPTED, REJECTED, CH_MAP, part_matcher, false)
	}

	// Check if workflow named 'in' exists
	in_ch, ok := CH_MAP["in"]
	if !ok {
		return -1, fmt.Errorf("no workflow named 'in'")
	}

	accepted := make([]Part, 0)
	rejected := make([]Part, 0)
	done := make(chan bool)
	go collect_results(1*time.Millisecond, done, ACCEPTED, REJECTED, &accepted, &rejected)

	// Send parts to workflow servers
	send_all(parts, in_ch, 10*time.Microsecond)

	// Wait for all parts to be processed
	<-done

	number_of_parts := len(accepted) + len(rejected)

	if number_of_parts != len(parts) {
		return -1, fmt.Errorf("not all parts were processed. Expected %d, got %d", len(parts), number_of_parts)
	}

	ratings := 0
	for _, part := range accepted {
		ratings += part.X
		ratings += part.M
		ratings += part.A
		ratings += part.S
	}

	return ratings, nil
}

type MatchingFunc[T any] func(T, Rule) ([]T, []T, error)

func workflow_server[T any](
	w Workflow,
	in chan T,
	accepted chan T,
	rejected chan T,
	chmap map[string]chan T,
	matching_func MatchingFunc[T],
	verbose bool,
) {
	for p := range in {
		if verbose {
			fmt.Printf("Workflow '%s' processing part %v\n", w.name, p)
		}

		// Parts which are lest for processing
		var parts []T = []T{p}
		for _, rule := range w.rules {
			if len(parts) == 0 {
				// No more parts to process
				break
			}
			// Pop the first part
			part := parts[0]
			parts = parts[1:]

			// Run the matching function
			matched_parts, unmatched_parts, err := matching_func(part, rule)

			if err != nil {
				fmt.Println(err)
				break
			}

			// Process the results and keep track of unmatched parts
			for _, part := range matched_parts {
				switch rule.target_type {
				case TGT_ACCEPT:
					accepted <- part
				case TGT_REJECT:
					rejected <- part
				case TGT_NAME:
					ch, ok := chmap[rule.target]
					if !ok {
					} else {
						ch <- part
					}
				}
			}

			// Update the parts list
			parts = append(parts, unmatched_parts...)
		}
	}
	if verbose {
		fmt.Printf("Workflow '%s' done\n", w.name)
	}
}

func part_matcher(p Part, r Rule) ([]Part, []Part, error) {
	ok, err := p.Matches(r)
	if err != nil {
		return nil, nil, err
	}
	if ok {
		return []Part{p}, nil, nil
	} else {
		return nil, []Part{p}, nil
	}
}

func (p Part) Matches(r Rule) (bool, error) {
	if r.op == OP_UNCOND {
		return true, nil
	} else if r.op == OP_GT {
		switch r.field {
		case FLD_X:
			return p.X > r.value, nil
		case FLD_M:
			return p.M > r.value, nil
		case FLD_A:
			return p.A > r.value, nil
		case FLD_S:
			return p.S > r.value, nil
		default:
			return false, fmt.Errorf("invalid field: %s", r.field)
		}
	} else if r.op == OP_LT {
		switch r.field {
		case FLD_X:
			return p.X < r.value, nil
		case FLD_M:
			return p.M < r.value, nil
		case FLD_A:
			return p.A < r.value, nil
		case FLD_S:
			return p.S < r.value, nil
		default:
			return false, fmt.Errorf("invalid field: %s", r.field)
		}
	} else {
		return false, fmt.Errorf("invalid op: %s", r.op)
	}
}

func collect_results[T any](
	timeout time.Duration,
	done chan bool,
	accepted chan T,
	rejected chan T,
	accepted_out *[]T,
	rejected_out *[]T,
) {
	last_activity := time.Now()
	for {
		select {
		case p := <-accepted:
			*accepted_out = append(*accepted_out, p)
			last_activity = time.Now()
		case p := <-rejected:
			*rejected_out = append(*rejected_out, p)
			last_activity = time.Now()
		default:
			if time.Since(last_activity) > timeout {
				done <- true
				return
			}
		}
	}
}

func send_all(parts []Part, ch chan Part, sleep time.Duration) {
	for _, part := range parts {
	retry:
		select {
		case ch <- part:
			// fmt.Println("Sent part", part)
		default:
			time.Sleep(sleep)
			goto retry
		}
	}
	// fmt.Println("All parts sent")
}
