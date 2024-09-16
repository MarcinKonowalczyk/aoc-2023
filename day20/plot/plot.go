package main

import (
	"aoc2023/day20"
	"aoc2023/utils"
	"flag"
	"fmt"
	"os"
	"os/exec"
)

var file string
var verbose bool

func init() {
	flag.Usage = func() {
		fmt.Println("Usage: go run ./day20/plot -filename <filename> [-v]")
		flag.PrintDefaults()
		os.Exit(1)
	}
	flag.StringVar(&file, "filename", "", "Input filename")
	flag.BoolVar(&verbose, "v", false, "Verbose output")
}

func main() {
	flag.Parse()

	if file == "" {
		flag.Usage()
	}

	fmt.Println("Plotting", file)

	lines, err := utils.FileToLines(file)
	if err != nil {
		utils.Stopf("Error when reading file. Error: %s", err)
	}

	modules, _, err := day20.ParseModules(lines, false)
	if err != nil {
		utils.Stopf("Error when parsing modules. Error: %s", err)
	}

	// find rx
	rx, ok := modules["rx"]
	if !ok {
		utils.Stopf("Module 'rx' not found")
	}

	if len(rx.Inputs()) != 1 {
		utils.Stopf("Module 'rx' should have exactly one input")
	}

	rxi := rx.Inputs()[0]

	if len(modules[rxi].Inputs()) != len(modules["broadcaster"].Targets()) {
		utils.Stopf("Module '%s' should have %d inputs", rxi, len(modules["broadcaster"].Targets()))
	}

	if verbose {
		fmt.Printf("rx's ancestor is '%s'. There are %d subgraphs.\n", rxi, len(modules[rxi].Inputs()))
	}

	// Open a new file
	f, err := os.Create("plot.dot")
	if err != nil {
		utils.Stopf("Error when creating file. Error: %s", err)
	}
	defer f.Close()

	// Write the header
	f.WriteString("digraph G {\n")
	f.WriteString("  rankdir=TB;\n")

	// Write the nodes
	for name, module := range modules {

		shape := kindToShape(module.Kind())
		f.WriteString(fmt.Sprintf("  %s [shape=%s];\n", name, shape))
	}

	// Write the edges
	for name, module := range modules {
		targets := module.Targets()
		color := kindToColor(module.Kind())
		for _, target := range targets {
			f.WriteString(fmt.Sprintf("  %s -> %s [color=%s];\n", name, target, color))
		}
	}

	// Write the footer
	f.WriteString("}\n")

	// Check if we have the dot command
	_, err = exec.LookPath("dot")
	if err != nil {
		fmt.Println("Please install Graphviz to generate the plot")
	} else {
		args := []string{"plot.dot", "-Tpng", "-o", "plot.png"}
		cmd := exec.Command("dot", args...)

		err = cmd.Run()
		if err != nil {
			utils.Stopf("Error when running command. Error: %s", err)
		}
	}
}

func kindToShape(kind day20.ModuleKind) string {
	switch kind {
	case day20.CONJUNCTION:
		return "diamond"
	case day20.BROADCASTER:
		return "ellipse"
	case day20.DUMMY:
		return "ellipse"
	}
	return "box"
}

func kindToColor(kind day20.ModuleKind) string {
	switch kind {
	case day20.FLIP_FLOP:
		return "red"
	case day20.CONJUNCTION:
		return "blue"
	}
	return "black"
}
