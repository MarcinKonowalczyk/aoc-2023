package day19

import (
	"fmt"
	"time"
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

func main_1(lines []string) (n int, err error) {
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

	// fmt.Println("Workflows:")
	// for _, workflow := range workflows {
	// 	fmt.Println(" ", workflow)
	// }

	// fmt.Println("Parts:")
	// for _, part := range parts {
	// 	fmt.Println(" ", part)
	// }

	// Create part input channels
	CH_MAP = make(map[string]chan Part)
	for _, workflow := range workflows {
		CH_MAP[workflow.name] = make(chan Part, MAX_PARTS)
	}

	REJECTED = make(chan Part, MAX_PARTS)
	ACCEPTED = make(chan Part, MAX_PARTS)

	// Start workflow servers
	for _, workflow := range workflows {
		go workflow_server(workflow, CH_MAP[workflow.name])
	}

	// Check if workflow named 'in' exists
	in_ch, ok := CH_MAP["in"]
	if !ok {
		return -1, fmt.Errorf("no workflow named 'in'")
	}

	accepted := make([]Part, 0)
	rejected := make([]Part, 0)
	done := make(chan bool)
	go collect_results(1*time.Millisecond, done, &accepted, &rejected)

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

func workflow_server(w Workflow, in chan Part) {
	for p := range in {
		for _, rule := range w.rules {
			ok, err := rule.Matches(p)
			if err != nil {
			}
			if ok {
				switch rule.target_type {
				case TGT_ACCEPT:
					ACCEPTED <- p
				case TGT_REJECT:
					REJECTED <- p
				case TGT_NAME:
					ch, ok := CH_MAP[rule.target]
					if !ok {
					} else {
						ch <- p
					}
				}
				break
			}
		}
	}
	fmt.Printf("Workflow '%s' done\n", w.name)
}

func (r Rule) Matches(p Part) (bool, error) {
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

func collect_results(timeout time.Duration, done chan bool, accepted *[]Part, rejected *[]Part) {
	last_activity := time.Now()
	for {
		select {
		case p := <-ACCEPTED:
			*accepted = append(*accepted, p)
			last_activity = time.Now()
		case p := <-REJECTED:
			*rejected = append(*rejected, p)
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
