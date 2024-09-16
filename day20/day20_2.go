package day20

import "fmt"

func main_2(lines []string, verbose bool) (n int, err error) {
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
	N_TEST_PULSES := 10000
	lows := make([]int, N_TEST_PULSES)
	highs := make([]int, N_TEST_PULSES)

	hooks := make(map[string]Hook)

	rx_called_with_low := false
	hooks["rx"] = func(sender string, value bool) {
		// fmt.Println("rx", sender, value)
		if !value {
			fmt.Println("rx called with low!!")
			rx_called_with_low = true
		}
	}
	fmt.Println("rx_called_with_low:", rx_called_with_low)

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
