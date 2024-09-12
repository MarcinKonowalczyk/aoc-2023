package day19

import (
	"fmt"
	"time"
)

var QCH_MAP map[string]chan QuantumPart
var QACCEPTED chan QuantumPart
var QREJECTED chan QuantumPart

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

	// Create part input channels
	QCH_MAP = make(map[string]chan QuantumPart)
	for _, workflow := range workflows {
		QCH_MAP[workflow.name] = make(chan QuantumPart, MAX_PARTS)
	}

	QACCEPTED = make(chan QuantumPart, MAX_PARTS)
	QREJECTED = make(chan QuantumPart, MAX_PARTS)

	// Start workflow servers
	for _, workflow := range workflows {
		ch := QCH_MAP[workflow.name]
		go workflow_server(workflow, ch, QACCEPTED, QREJECTED, QCH_MAP, qpart_matcher, false)
	}

	// Check if workflow named 'in' exists
	in_ch, ok := QCH_MAP["in"]
	if !ok {
		return -1, fmt.Errorf("no workflow named 'in'")
	}

	accepted := make([]QuantumPart, 0)
	rejected := make([]QuantumPart, 0)
	done := make(chan bool)
	go collect_results(1*time.Millisecond, done, QACCEPTED, QREJECTED, &accepted, &rejected)

	// Send parts to workflow servers
	in_ch <- NewQuantumPart()

	// Wait for all parts to be processed
	<-done

	number_of_parts := len(accepted) + len(rejected)

	if verbose {
		fmt.Println("Number of parts:", number_of_parts)
		fmt.Println("Accepted parts:")
		for _, part := range accepted {
			fmt.Println(" ", part)
		}
		fmt.Println("Rejected parts:")
		for _, part := range rejected {
			fmt.Println(" ", part)
		}
	}

	total_number_of_states := 0
	for _, part := range accepted {
		total_number_of_states += part.NumberOfStates()
	}

	return total_number_of_states, nil
}

type QuantumPart struct {
	Xmin int
	Xmax int
	Mmin int
	Mmax int
	Amin int
	Amax int
	Smin int
	Smax int
}

func (p QuantumPart) String() string {
	return fmt.Sprintf("Q{x=%d|%d, m=%d|%d, a=%d|%d, s=%d|%d}", p.Xmin, p.Xmax, p.Mmin, p.Mmax, p.Amin, p.Amax, p.Smin, p.Smax)
}

func NewQuantumPart() QuantumPart {
	return QuantumPart{
		Xmin: MIN_PART_VALUE,
		Xmax: MAX_PART_VALUE,
		Mmin: MIN_PART_VALUE,
		Mmax: MAX_PART_VALUE,
		Amin: MIN_PART_VALUE,
		Amax: MAX_PART_VALUE,
		Smin: MIN_PART_VALUE,
		Smax: MAX_PART_VALUE,
	}
}

func (p QuantumPart) Copy() QuantumPart {
	return QuantumPart{
		Xmin: p.Xmin,
		Xmax: p.Xmax,
		Mmin: p.Mmin,
		Mmax: p.Mmax,
		Amin: p.Amin,
		Amax: p.Amax,
		Smin: p.Smin,
		Smax: p.Smax,
	}
}

func qpart_matcher(p QuantumPart, r Rule) ([]QuantumPart, []QuantumPart, error) {
	matched, unmatched, err := p.QuantumMatches(r)
	if err != nil {
		return nil, nil, err
	}
	if matched.Xmin > matched.Xmax || matched.Mmin > matched.Mmax || matched.Amin > matched.Amax || matched.Smin > matched.Smax {
		return nil, nil, fmt.Errorf("invalid matching result: %v", matched)
	}
	if unmatched.Xmin > unmatched.Xmax || unmatched.Mmin > unmatched.Mmax || unmatched.Amin > unmatched.Amax || unmatched.Smin > unmatched.Smax {
		return nil, nil, fmt.Errorf("invalid matching result: %v", unmatched)
	}
	return []QuantumPart{matched}, []QuantumPart{unmatched}, nil
}

func (p QuantumPart) QuantumMatches(r Rule) (QuantumPart, QuantumPart, error) {
	if r.op == OP_UNCOND {
		return p, QuantumPart{}, nil
	} else if r.op == OP_GT {
		switch r.field {
		case FLD_X:
			if p.Xmin > r.value {
				return p, QuantumPart{}, nil
			} else if p.Xmax < r.value {
				return QuantumPart{}, p, nil
			} else {
				mp := p.Copy()
				mp.Xmin = r.value + 1
				np := p.Copy()
				np.Xmax = r.value
				return mp, np, nil
			}
		case FLD_M:
			if p.Mmin > r.value {
				return p, QuantumPart{}, nil
			} else if p.Mmax < r.value {
				return QuantumPart{}, p, nil
			} else {
				mp := p.Copy()
				mp.Mmin = r.value + 1
				np := p.Copy()
				np.Mmax = r.value
				return mp, np, nil
			}
		case FLD_A:
			if p.Amin > r.value {
				return p, QuantumPart{}, nil
			} else if p.Amax < r.value {
				return QuantumPart{}, p, nil
			} else {
				mp := p.Copy()
				mp.Amin = r.value + 1
				np := p.Copy()
				np.Amax = r.value
				return mp, np, nil
			}
		case FLD_S:
			if p.Smin > r.value {
				return p, QuantumPart{}, nil
			} else if p.Smax < r.value {
				return QuantumPart{}, p, nil
			} else {
				mp := p.Copy()
				mp.Smin = r.value + 1
				np := p.Copy()
				np.Smax = r.value
				return mp, np, nil
			}
		default:
			return QuantumPart{}, QuantumPart{}, fmt.Errorf("invalid field: %s", r.field)
		}
	} else if r.op == OP_LT {
		switch r.field {
		case FLD_X:
			if p.Xmin > r.value {
				return QuantumPart{}, p, nil
			} else if p.Xmax < r.value {
				return p, QuantumPart{}, nil
			} else {
				mp := p.Copy()
				mp.Xmax = r.value - 1
				np := p.Copy()
				np.Xmin = r.value
				return mp, np, nil
			}
		case FLD_M:
			if p.Mmin > r.value {
				return QuantumPart{}, p, nil
			} else if p.Mmax < r.value {
				return p, QuantumPart{}, nil
			} else {
				mp := p.Copy()
				mp.Mmax = r.value - 1
				np := p.Copy()
				np.Mmin = r.value
				return mp, np, nil
			}
		case FLD_A:
			if p.Amin > r.value {
				return QuantumPart{}, p, nil
			} else if p.Amax < r.value {
				return p, QuantumPart{}, nil
			} else {
				mp := p.Copy()
				mp.Amax = r.value - 1
				np := p.Copy()
				np.Amin = r.value
				return mp, np, nil
			}
		case FLD_S:
			if p.Smin > r.value {
				return QuantumPart{}, p, nil
			} else if p.Smax < r.value {
				return p, QuantumPart{}, nil
			} else {
				mp := p.Copy()
				mp.Smax = r.value - 1
				np := p.Copy()
				np.Smin = r.value
				return mp, np, nil
			}
		default:
			return QuantumPart{}, QuantumPart{}, fmt.Errorf("invalid field: %s", r.field)
		}
	} else {
		return QuantumPart{}, QuantumPart{}, fmt.Errorf("invalid op: %s", r.op)
	}
}

func (p QuantumPart) NumberOfStates() int {
	return (p.Xmax - p.Xmin + 1) * (p.Mmax - p.Mmin + 1) * (p.Amax - p.Amin + 1) * (p.Smax - p.Smin + 1)
}
