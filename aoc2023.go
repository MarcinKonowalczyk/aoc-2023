package main

import (
	"aoc2023/day01"
	"aoc2023/day02"
	"aoc2023/day03"
	"aoc2023/day04"
	"aoc2023/day05"
	"aoc2023/day06"
	"aoc2023/day07"
	"aoc2023/day08"
	"aoc2023/day09"
	"aoc2023/day10"
	"aoc2023/day11"
	"aoc2023/day12"
	"aoc2023/day13"
	"aoc2023/day14"
	"aoc2023/day15"
	"aoc2023/day16"
	"aoc2023/day17"
	"aoc2023/day18"
	"aoc2023/day19"
	"aoc2023/day20"
	"aoc2023/day21"
	"aoc2023/utils"
	"flag"
	"fmt"
	"os"
	"time"
)

var day int
var part int
var file string
var verbose bool

func init() {
	flag.Usage = func() {
		fmt.Println("Usage: aoc2023 -day <day> -part <part> -filename <filename> [-v]")
		flag.PrintDefaults()
		os.Exit(1)
	}
	flag.IntVar(&day, "day", 0, "Day of the Advent of Code 2023")
	flag.IntVar(&part, "part", 0, "Part of the task")
	flag.StringVar(&file, "filename", "", "Input filename")
	flag.BoolVar(&verbose, "v", false, "Verbose output")
}

func main() {
	flag.Parse()

	if day == 0 || part == 0 || file == "" {
		flag.Usage()
	}

	if day < 1 || day > 25 {
		utils.Stopf("Expected day in rage [0, 25]. Got: %d", day)
	}

	if part < 1 || part > 2 {
		utils.Stopf("Expected part in rage [1, 2]. Got: %d", part)
	}

	// Resolve path
	file, err := utils.ResolvePath(file)
	if err != nil {
		utils.Stopf("Error when resolving path. Error: %s", err)
	}

	fmt.Println("Running AOC 2023 code")
	fmt.Println("Got day:", day)
	fmt.Println("Got part:", part)
	fmt.Println("Got file:", file)

	// Read input file
	lines, err := utils.FileToLines(file)
	if err != nil {
		utils.Stopf("Error when reading file. Error: %s", err)
	}

	// If the last line is empty, remove it
	if len(lines) > 0 && len(lines[len(lines)-1]) == 0 {
		lines = lines[:len(lines)-1]
	}

	var value int
	tic := time.Now()

	defer func() {
		if r := recover(); r != nil {
			utils.Stopf("Recovered from panic in main: %v", r)
		}
	}()

	switch day {
	case 1:
		value, err = day01.Main(part, lines, verbose)
	case 2:
		value, err = day02.Main(part, lines, verbose)
	case 3:
		value, err = day03.Main(part, lines, verbose)
	case 4:
		value, err = day04.Main(part, lines, verbose)
	case 5:
		value, err = day05.Main(part, lines, verbose)
	case 6:
		value, err = day06.Main(part, lines, verbose)
	case 7:
		value, err = day07.Main(part, lines, verbose)
	case 8:
		value, err = day08.Main(part, lines, verbose)
	case 9:
		value, err = day09.Main(part, lines, verbose)
	case 10:
		value, err = day10.Main(part, lines, verbose)
	case 11:
		value, err = day11.Main(part, lines, verbose)
	case 12:
		value, err = day12.Main(part, lines, verbose)
	case 13:
		value, err = day13.Main(part, lines, verbose)
	case 14:
		value, err = day14.Main(part, lines, verbose)
	case 15:
		value, err = day15.Main(part, lines, verbose)
	case 16:
		value, err = day16.Main(part, lines, verbose)
	case 17:
		value, err = day17.Main(part, lines, verbose)
	case 18:
		value, err = day18.Main(part, lines, verbose)
	case 19:
		value, err = day19.Main(part, lines, verbose)
	case 20:
		value, err = day20.Main(part, lines, verbose)
	case 21:
		value, err = day21.Main(part, lines, verbose)
	// case 22:
	// 	value, err = day22.Main(part, lines, verbose)
	// case 23:
	// 	value, err = day23.Main(part, lines, verbose)
	// case 24:
	// 	value, err = day24.Main(part, lines, verbose)
	// case 25:
	// 	value, err = day25.Main(part, lines, verbose)
	default:
		utils.Stopf("Day %d is not implemented yet", day)
	}
	toc := time.Now()

	if err != nil {
		utils.Stopf("Error when running main. Error: %s", err)
	}

	fmt.Printf("Return value: %d\n", value)
	fmt.Printf("Execution time: %v\n", toc.Sub(tic))

	_ = utils.CopyToClipboard(fmt.Sprintf("%d", value))

}
