package day05

import (
	"aoc2023/utils"
	"fmt"
	"regexp"
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

type MapLine struct {
	source_start      int
	destination_start int
	length            int
}

type Map []MapLine

func lookUpValueInMapLine(value int, map_line MapLine) int {
	source_start := map_line.source_start
	source_end := map_line.source_start + map_line.length
	map_delta := map_line.destination_start - map_line.source_start
	if value >= source_start && value <= source_end {
		return value + map_delta
	}
	return value
}

func lookUpValueInMap(value int, map_lines []MapLine) int {
	for _, map_line := range map_lines {
		mapped_value := lookUpValueInMapLine(value, map_line)
		if mapped_value != value {
			return mapped_value
		}
	}
	return value
}

func main_1(lines []string) (n int, err error) {
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

	// fmt.Println("seeds:", seeds)
	// fmt.Println("seed_to_soil_map:", seed_to_soil_map)
	// fmt.Println("soil_to_fertilizer_map:", soil_to_fertilizer_map)
	// fmt.Println("ferilizer_to_water_map:", ferilizer_to_water_map)
	// fmt.Println("water_to_light_map:", water_to_light_map)
	// fmt.Println("light_to_temperature_map:", light_to_temperature_map)
	// fmt.Println("temperature_to_humidity_map:", temperature_to_humidity_map)
	// fmt.Println("humidity_to_location_map:", humidity_to_location_map)

	var lowest_location int
	for seed_index, seed := range seeds {
		soil := lookUpValueInMap(seed, seed_to_soil_map)
		fertilizer := lookUpValueInMap(soil, soil_to_fertilizer_map)
		water := lookUpValueInMap(fertilizer, ferilizer_to_water_map)
		light := lookUpValueInMap(water, water_to_light_map)
		temperature := lookUpValueInMap(light, light_to_temperature_map)
		humidity := lookUpValueInMap(temperature, temperature_to_humidity_map)
		location := lookUpValueInMap(humidity, humidity_to_location_map)
		// fmt.Printf("seed %d: %d -> %d -> %d -> %d -> %d -> %d -> %d -> %d\n", seed_index, seed, soil, fertilizer, water, light, temperature, humidity, location)
		if seed_index == 0 || location < lowest_location {
			lowest_location = location
		}
	}

	return lowest_location, nil
}

func parseMapFromLines(lines []string, header string) ([]MapLine, []string, error) {
	_, lines, err := cutAtLine(lines, header+" map:")
	if err != nil {
		return nil, nil, err
	}
	map_lines, lines, _ := cutAtLine(lines, "")
	// if err != nil {
	// 	return nil, nil, err
	// }
	map_lines = utils.ArrayFilter(map_lines, func(line string) bool {
		return line != ""
	})

	parsed_map, err := utils.ArrayMapWithError(map_lines, func(line string) (MapLine, error) {
		result, err := utils.StringOfNumbersToInts(line)
		if err != nil {
			return MapLine{}, err
		}
		if len(result) != 3 {
			return MapLine{}, fmt.Errorf("invalid seed-to-soil map line: %s", line)
		}
		return MapLine{
			source_start:      result[1],
			destination_start: result[0],
			length:            result[2],
		}, nil
	})
	if err != nil {
		return nil, nil, err
	}

	return parsed_map, lines, nil
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

var SEEDS_LINE_RE = regexp.MustCompile(`seeds: (?P<seeds>(\d+ *)+)`)

func parseSeedsLine(line string) ([]int, error) {
	matches := utils.GetNamedSubexpsCompiledRe(SEEDS_LINE_RE, line)
	seeds, err := utils.StringOfNumbersToInts(matches["seeds"])
	if err != nil {
		return nil, err
	}
	if len(seeds)%2 != 0 {
		return nil, fmt.Errorf("number of seeds must be even")
	}
	return seeds, nil
}
