package day05

type valueRange struct {
	start  int
	length int
}

func main_2(lines []string) (n int, err error) {
	seeds, err := parseSeedsLine(lines[0])
	if err != nil {
		return -1, err
	}
	seed_to_soil_map, lines, err := parseMapFromLines(lines, "seed-to-soil")
	if err != nil {
		return -1, err
	}
	soil_to_fertilizer_map, lines, err := parseMapFromLines(lines, "soil-to-fertilizer")
	if err != nil {
		return -1, err
	}
	ferilizer_to_water_map, lines, err := parseMapFromLines(lines, "fertilizer-to-water")
	if err != nil {
		return -1, err
	}
	water_to_light_map, lines, err := parseMapFromLines(lines, "water-to-light")
	if err != nil {
		return -1, err
	}
	light_to_temperature_map, lines, err := parseMapFromLines(lines, "light-to-temperature")
	if err != nil {
		return -1, err
	}
	temperature_to_humidity_map, lines, err := parseMapFromLines(lines, "temperature-to-humidity")
	if err != nil {
		return -1, err
	}
	humidity_to_location_map, _, err := parseMapFromLines(lines, "humidity-to-location")
	if err != nil {
		return -1, err
	}

	seed_ranges := []valueRange{}
	for i := 0; i < len(seeds); i += 2 {
		seed_ranges = append(seed_ranges, valueRange{start: seeds[i], length: seeds[i+1]})
	}

	chain_of_maps := []Map{
		seed_to_soil_map,
		soil_to_fertilizer_map,
		ferilizer_to_water_map,
		water_to_light_map,
		light_to_temperature_map,
		temperature_to_humidity_map,
		humidity_to_location_map,
	}

	all_location_ranges := []valueRange{}
	for _, seed_range := range seed_ranges {
		location_ranges := lookUpValueRangeInChainOfMaps(seed_range, chain_of_maps)
		all_location_ranges = append(all_location_ranges, location_ranges...)
	}
	lowest_location := -1
	for _, location_range := range all_location_ranges {
		if lowest_location == -1 || location_range.start < lowest_location {
			lowest_location = location_range.start
		}
	}

	return lowest_location, nil
}

func lookUpValueRangeInMapLine(value_range valueRange, map_line MapLine) ([]valueRange, []valueRange) {
	mapped_to_themselves := make([]valueRange, 0)
	mapped_through := make([]valueRange, 0)
	value_range_start := value_range.start
	value_range_end := value_range.start + value_range.length
	map_line_start := map_line.source_start
	map_line_end := map_line.source_start + map_line.length
	map_delta := map_line.destination_start - map_line.source_start
	if value_range_start < map_line_start {
		if value_range_end < map_line_start {
			// Value range is completely before map line. It maps to itself.
			mapped_to_themselves = append(mapped_to_themselves, value_range)
		} else if value_range_end <= map_line_end {
			// Value range starts before map line and ends within map line.
			// The bit before the map line maps to itself.
			mapped_to_themselves = append(mapped_to_themselves, valueRange{start: value_range_start, length: map_line_start - value_range_start})
			// The bit within the map line maps through it.
			mapped_through = append(mapped_through, valueRange{start: map_line_start + map_delta, length: value_range_end - map_line_start})
		} else {
			// Value range starts before map line and ends after map line.
			// The bit before the map line maps to itself.
			mapped_to_themselves = append(mapped_to_themselves, valueRange{start: value_range_start, length: map_line_start - value_range_start})
			// The bit within the map line maps through it.
			mapped_through = append(mapped_through, valueRange{start: map_line_start + map_delta, length: map_line.length})
			// The bit after the map line maps to itself.
			mapped_to_themselves = append(mapped_to_themselves, valueRange{start: map_line_end, length: value_range_end - map_line_end})
		}
	} else if value_range_start <= map_line_end {
		if value_range_end <= map_line_end {
			// Value range is completely within map line.
			mapped_through = append(mapped_through, valueRange{start: value_range_start + map_delta, length: value_range.length})
		} else {
			// Value range starts within map line and ends after map line.
			// The bit within the map line maps through it.
			mapped_through = append(mapped_through, valueRange{start: value_range_start + map_delta, length: map_line_end - value_range_start})
			// The bit after the map line maps to itself.
			mapped_to_themselves = append(mapped_to_themselves, valueRange{start: map_line_end, length: value_range_end - map_line_end})
		}
	} else {
		// Value range is completely after map line. It maps to itself.
		mapped_to_themselves = append(mapped_to_themselves, value_range)
	}

	return mapped_to_themselves, mapped_through
}

// Look up a value range in a map. The map can output multiple value ranges.
func lookUpValueRangeInMap(value_range valueRange, m Map) []valueRange {
	var out []valueRange = []valueRange{}
	r := []valueRange{value_range}
	for _, map_line := range m {
		nr := make([]valueRange, 0)
		for _, range_to_look_up := range r {
			mapped_to_themselves, mapped_through := lookUpValueRangeInMapLine(range_to_look_up, map_line)
			// The bits which got mapped through can go straight into the output.
			out = append(out, mapped_through...)
			// The other bits need to be looked up in other map lines.
			nr = append(nr, mapped_to_themselves...)
		}
		r = nr
	}
	// These are the bits which got mapped to themselves, even after all the map lines.
	// They are supposed to map to themselves, so add them to the output.
	out = append(out, r...)
	return out
}

// Look up a value range in a chain of maps. For each output of the first map,
// it is then looked up in the second map, and so on.
func lookUpValueRangeInChainOfMaps(value_range valueRange, maps []Map) []valueRange {
	r := []valueRange{value_range}
	for _, m := range maps {
		// Each map can output multiple value ranges, so for each map we need to
		// collect them all into a new slice and then use that as the input for
		// the next map.
		nr := make([]valueRange, 0)
		for _, range_to_look_up := range r {
			mapped := lookUpValueRangeInMap(range_to_look_up, m)
			nr = append(nr, mapped...)
		}
		r = nr
	}
	return r
}
