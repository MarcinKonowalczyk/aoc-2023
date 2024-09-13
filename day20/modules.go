package day20

import (
	"aoc2023/utils"
	"fmt"
	"strings"
)

type Output int8

const (
	NOTHING Output = -1
	LOW     Output = 0
	HIGH    Output = 1
)

func (o Output) String() string {
	switch o {
	case NOTHING:
		return "nothing"
	case LOW:
		return "low"
	case HIGH:
		return "high"
	default:
		panic("invalid Output")
	}
}
func (o Output) Bool() bool {
	if o == NOTHING {
		panic("Output.NOTHING has no boolean value")
	}
	return o == HIGH
}

type Module interface {
	String() string
	Name() string
	Targets() []string
	Inputs() []string
	AddInput(string)
	Send(string, bool) Output
}

////////// FlipFlop //////////

type FlipFlop struct {
	name    string
	targets []string
	inputs  []string
	state   bool
}

func NewFlipFlop(name string, targets []string) *FlipFlop {
	return &FlipFlop{name: name, targets: targets, state: false}
}

func (f *FlipFlop) String() string {
	state := utils.Ternary(f.state, "1", "0")
	return fmt.Sprintf(
		"%%%s(%s) -> %v (%v)",
		f.name,
		state,
		f.targets,
		f.inputs,
	)
}

func (f *FlipFlop) Name() string {
	return f.name
}

func (f *FlipFlop) Targets() []string {
	return f.targets
}

func (f *FlipFlop) Inputs() []string {
	return f.inputs
}

func (f *FlipFlop) AddInput(input string) {
	f.inputs = append(f.inputs, input)
}

func (f *FlipFlop) Send(sender string, value bool) Output {
	// NOTE: FlipFlop ignores the sender
	if value {
		// High pulse. Nothing happens
		return NOTHING
	} else {
		// Low pulse. Flip and send state
		f.state = !f.state
		if f.state {
			return HIGH
		} else {
			return LOW
		}
	}

}

// Check that FlipFlop implements Module
var _ Module = &FlipFlop{}

////////// Conjunction //////////

type Conjunction struct {
	name     string
	targets  []string
	inputs   []string
	previous map[string]bool
}

func (c *Conjunction) String() string {
	previous_str := make([]string, 0)
	for _, name := range c.inputs {
		state := utils.Ternary(c.previous[name], "1", "0")
		previous_str = append(previous_str, fmt.Sprintf("%s=%v", name, state))
	}

	return fmt.Sprintf("&%s(%s) -> %v (%v)", c.name, previous_str, c.targets, c.inputs)
}

func (c *Conjunction) Name() string {
	return c.name
}

func (c *Conjunction) Targets() []string {
	return c.targets
}

func (c *Conjunction) Inputs() []string {
	return c.inputs
}

func (c *Conjunction) AddInput(input string) {
	c.inputs = append(c.inputs, input)
	if c.previous == nil {
		c.previous = make(map[string]bool)
	}
	c.previous[input] = false
}

func (c *Conjunction) Send(sender string, value bool) Output {
	// Update previous value
	c.previous[sender] = value
	// Check if all previous values are high
	all_high := true
	for _, v := range c.previous {
		if !v {
			all_high = false
			break
		}
	}
	// NOTE: Conjunction never returns NOTHING
	if all_high {
		return LOW
	} else {
		return HIGH
	}
}

// Check that Conjunction implements Module
var _ Module = &Conjunction{}

////////// Broadcaster //////////

type Broadcaster struct {
	targets []string
}

func (b *Broadcaster) String() string {
	return fmt.Sprintf("broadcaster -> %v", b.targets)
}

func (b *Broadcaster) Name() string {
	return "broadcaster"
}

func (b *Broadcaster) Targets() []string {
	return b.targets
}

func (b *Broadcaster) Inputs() []string {
	return nil
}

func (b *Broadcaster) AddInput(input string) {
	panic("broadcaster should not have inputs")
}

func (b *Broadcaster) Send(sender string, value bool) Output {
	// NOTE: Broadcaster ignores the sender
	if value {
		return HIGH
	} else {
		return LOW
	}
}

// Check that Broadcaster implements Module
var _ Module = &Broadcaster{}

////////// Dummy //////////

type Dummy struct {
	name   string
	inputs []string
}

func (d *Dummy) String() string {
	return fmt.Sprintf("%s (%v)", d.name, d.inputs)
}

func (d *Dummy) Name() string {
	return d.name
}

func (d *Dummy) Targets() []string {
	return nil
}

func (d *Dummy) Inputs() []string {
	return d.inputs
}

func (d *Dummy) AddInput(input string) {
	d.inputs = append(d.inputs, input)
}

func (d *Dummy) Send(sender string, value bool) Output {
	return NOTHING
}

func NewDummy(name string) *Dummy {
	return &Dummy{name: name}
}

// Check that Dummy implements Module
var _ Module = &Dummy{}

////////// Parse //////////

func ParseLine(line string) (Module, error) {
	parts := strings.Split(line, " -> ")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid line: %s", line)
	}
	targets := strings.Split(parts[1], ", ")
	if parts[0] == "broadcaster" {
		return &Broadcaster{targets: targets}, nil
	} else {
		name := parts[0]
		if name[0] == '%' {
			return NewFlipFlop(name[1:], targets), nil
		} else if name[0] == '&' {
			return &Conjunction{name: name[1:], targets: targets}, nil
		} else {
			return nil, fmt.Errorf("invalid line: %s", line)
		}
	}

}
