package day20

import (
	"aoc2023/utils"
	"fmt"
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

type element struct {
	sender string
	value  bool
}

func main_1(lines []string, verbose bool) (n int, err error) {
	modules, names, err := ParseModules(lines, verbose)
	if err != nil {
		return 0, err
	}

	// Make the input queues
	input_queues := make(map[string][]element, 0)
	for name := range modules {
		input_queues[name] = make([]element, 0)
	}

	// Send a bunch of test pulses to the broadcaster
	N_TEST_PULSES := 1000
	lows := make([]int, N_TEST_PULSES)
	highs := make([]int, N_TEST_PULSES)

	hooks := make(map[string]Hook)

	for i := 0; i < N_TEST_PULSES; i++ {
		input_queues["broadcaster"] = []element{{"", false}}
		n_low, n_high, err := drive_machine(modules, input_queues, names, hooks, false)
		if err != nil {
			return 0, err
		}
		n_low++ // Count the initial pulse
		lows[i] = n_low
		highs[i] = n_high
	}

	total_n_low := 0
	total_n_high := 0
	for i := 0; i < N_TEST_PULSES; i++ {
		total_n_low += lows[i]
		total_n_high += highs[i]
	}
	product := total_n_low * total_n_high

	return product, nil
}

func ParseModules(lines []string, verbose bool) (map[string]Module, []string, error) {
	modules := make(map[string]Module, 0)
	names := make([]string, 0) // Names in order of appearance
	for _, line := range lines {
		module, err := ParseLine(line)
		if err != nil {
			return nil, nil, err
		}
		name := module.Name()
		if _, ok := modules[name]; ok {
			return nil, nil, fmt.Errorf("duplicate module: %s", name)
		}
		modules[name] = module
		names = append(names, name)
	}

	// Make sure we have a single broadcaster
	found := utils.ArrayContainsFunc(names, func(name string) bool {
		return name == "broadcaster"
	})
	if !found {
		return nil, nil, fmt.Errorf("missing broadcaster")
	}

	// Link inputs to outputs
	missing_targets := make([]string, 0) // set of missing targets
	for _, module := range modules {
		targets := make([]Module, 0)
		for _, target_name := range module.Targets() {
			target, ok := modules[target_name]
			if !ok {
				missing_targets = append(missing_targets, target_name)
				dummy := NewDummy(target_name)
				modules[target_name] = dummy
				target = dummy
			}
			targets = append(targets, target)
		}
		name := module.Name()
		for _, target := range targets {
			target.AddInput(name)
		}
	}

	if verbose {
		fmt.Println("Missing targets:", missing_targets)
	}

	names = append(names, missing_targets...)

	// Print modules
	if verbose {
		fmt.Println("Modules:")
		for _, name := range names {
			module := modules[name]
			fmt.Printf(" %v\n", module)
		}
	}

	return modules, names, nil
}

type Hook func(sender string, value bool)

func drive_machine(
	modules map[string]Module,
	input_queues map[string][]element,
	names_in_order []string,
	module_hooks map[string]Hook,
	verbose bool,
) (int, int, error) {

	if len(modules) != len(input_queues) || len(modules) != len(names_in_order) {
		return -1, -1, fmt.Errorf("len(modules) (%d) != len(input_queues) (%d) != len(names_in_order) (%d)", len(modules), len(input_queues), len(names_in_order))
	}

	n_low := 0
	n_high := 0

	if len(modules) == 0 {
		return n_low, n_high, nil
	}

	for {
		// Find non-empty input queue
		all_empty := true
		var name string
		for _, name = range names_in_order {
			if len(input_queues[name]) > 0 {
				all_empty = false
				break
			}
		}

		if all_empty {
			// No more input
			break
		}

		// Pop first element
		e := input_queues[name][0]
		input_queues[name] = input_queues[name][1:]

		// Get the module
		module, ok := modules[name]
		if !ok {
			return -1, -1, fmt.Errorf("module not found: %s", name)
		}

		// Call the hook for this module if it exists
		if hook, ok := module_hooks[name]; ok {
			hook(e.sender, e.value)
		}
		output := module.Send(e.sender, e.value)
		if output == NOTHING {
			// Nothing to do. Pass
			continue
		}
		signal := output.Bool()

		// Send signal to targets
		targets := module.Targets()
		qe := element{module.Name(), signal}
		for _, target_name := range targets {
			if verbose {
				fmt.Print(name, " -", output, "-> ", target_name, "\n")
			}
			input_queues[target_name] = append(input_queues[target_name], qe)
		}

		// Count high and low signals
		if signal {
			n_high += len(targets)
		} else {
			n_low += len(targets)
		}
	}
	return n_low, n_high, nil
}
