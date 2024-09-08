package day19

import (
	"fmt"
	"strconv"
	"strings"
)

type Workflow struct {
	name  string
	rules []Rule
}

func (w Workflow) String() string {
	return fmt.Sprintf("%s: %v", w.name, w.rules)
}

type Op int

const (
	OP_LT Op = iota
	OP_GT
	OP_UNCOND
)

func (o Op) String() string {
	switch o {
	case OP_LT:
		return "<"
	case OP_GT:
		return ">"
	case OP_UNCOND:
		return ""
	default:
		return "UNKNOWN"
	}
}

type Field int

const (
	FLD_X Field = iota
	FLD_M
	FLD_A
	FLD_S
)

func (f Field) String() string {
	switch f {
	case FLD_X:
		return "x"
	case FLD_M:
		return "m"
	case FLD_A:
		return "a"
	case FLD_S:
		return "s"
	default:
		return "UNKNOWN"
	}
}

func ParseField(field string) (Field, error) {
	field = strings.ToLower(field)
	switch field {
	case "x":
		return FLD_X, nil
	case "m":
		return FLD_M, nil
	case "a":
		return FLD_A, nil
	case "s":
		return FLD_S, nil
	default:
		return FLD_X, fmt.Errorf("invalid field: %s", field)
	}
}

type TargetType int

const (
	TGT_REJECT TargetType = iota
	TGT_ACCEPT
	TGT_NAME
)

func (t TargetType) String() string {
	switch t {
	case TGT_REJECT:
		return "reject"
	case TGT_ACCEPT:
		return "accept"
	case TGT_NAME:
		return "name"
	default:
		return "UNKNOWN"
	}
}

type Rule struct {
	field       Field
	op          Op
	value       int
	target      string
	target_type TargetType
}

func (r Rule) String() string {
	if r.op == OP_UNCOND {
		return r.target
	}
	return fmt.Sprintf("%s%s%d:%s", r.field, r.op, r.value, r.target)
}

func ParseWorkflowLine(line string) (Workflow, error) {
	parts := strings.Split(line, "{")
	if len(parts) != 2 {
		return Workflow{}, fmt.Errorf("invalid rule line: %s", line)
	}
	name := parts[0]
	line = parts[1]
	if line[len(line)-1] != '}' {
		return Workflow{}, fmt.Errorf("invalid rule line: %s", line)
	}
	line = line[:len(line)-1]
	parts = strings.Split(line, ",")
	rules := make([]Rule, len(parts))
	for i, part := range parts {
		rule, err := ParseRulePart(part)
		if err != nil {
			return Workflow{}, err
		}
		rules[i] = rule
	}

	return Workflow{
		name:  name,
		rules: rules,
	}, nil
}

func ParseRulePart(part string) (Rule, error) {
	parts := strings.Split(part, ":")
	if len(parts) == 1 {
		return Rule{
			field:       FLD_X,
			op:          OP_UNCOND,
			value:       0,
			target:      part,
			target_type: ParseTargetType(part),
		}, nil
	} else if len(parts) != 2 {
		return Rule{}, fmt.Errorf("invalid rule part: %s", part)
	}
	// We've split at a singular colon
	target := parts[1]
	parts = strings.Split(parts[0], "<")
	if len(parts) == 2 {
		// We've split at a less than sign
		field, err := ParseField(parts[0])
		if err != nil {
			return Rule{}, err
		}
		value, err := strconv.Atoi(parts[1])
		if err != nil {
			return Rule{}, err
		}
		return Rule{
			field:       field,
			op:          OP_LT,
			value:       value,
			target:      target,
			target_type: ParseTargetType(target),
		}, nil
	}
	parts = strings.Split(parts[0], ">")
	if len(parts) == 2 {
		// We've split at a greater than sign
		field, err := ParseField(parts[0])
		if err != nil {
			return Rule{}, err
		}
		value, err := strconv.Atoi(parts[1])
		if err != nil {
			return Rule{}, err
		}
		return Rule{
			field:       field,
			op:          OP_GT,
			value:       value,
			target:      target,
			target_type: ParseTargetType(target),
		}, nil
	}
	// We failed to split at a less than or greater than sign
	return Rule{}, fmt.Errorf("invalid rule part: %s", part)
}

func ParseTargetType(target string) TargetType {
	target = strings.ToLower(target)
	if target == "r" {
		return TGT_REJECT
	} else if target == "a" {
		return TGT_ACCEPT
	} else {
		return TGT_NAME
	}
}

type Part struct {
	X int
	M int
	A int
	S int
}

func (p Part) String() string {
	return fmt.Sprintf("{x=%d m=%d a=%d s=%d}", p.X, p.M, p.A, p.S)
}
func ParsePartLine(line string) (Part, error) {
	if !strings.HasPrefix(line, "{") || !strings.HasSuffix(line, "}") {
		return Part{}, fmt.Errorf("invalid part line: %s", line)
	}
	line = line[1 : len(line)-1]
	parts := strings.Split(line, ",")
	if len(parts) != 4 {
		return Part{}, fmt.Errorf("invalid part line: %s", line)
	}
	p := Part{
		X: -1,
		M: -1,
		A: -1,
		S: -1,
	}
	for _, part := range parts {
		parts := strings.Split(part, "=")
		if len(parts) != 2 {
			return p, fmt.Errorf("invalid part line: %s", line)
		}
		value, err := strconv.Atoi(parts[1])
		if err != nil {
			return Part{}, err
		}
		if value < 0 {
			return Part{}, fmt.Errorf("invalid part line: %s", line)
		}
		switch parts[0] {
		case "x":
			p.X = value
		case "m":
			p.M = value
		case "a":
			p.A = value
		case "s":
			p.S = value
		default:
			return Part{}, fmt.Errorf("invalid part line: %s", line)
		}
	}

	if p.X == -1 || p.M == -1 || p.A == -1 || p.S == -1 {
		return Part{}, fmt.Errorf("invalid part line: %s", line)
	}
	return p, nil
}
