package main

import (
	"regexp"
	"strings"
)

func init() {
	RegisterSolver("day3", &Day3{})
}

type Day3 struct {
	input string
}

func (d *Day3) Init(input string) {
	d.input = strings.ReplaceAll(input, "\n", "")
}

func (d *Day3) Part1() int {
	return runMultInstructions(d.input)
}

func (d *Day3) Part2() int {
	trimmed := trimInstructions(d.input)
	return runMultInstructions(trimmed)
}

func trimInstructions(input string) string {
	re := regexp.MustCompile(`don't\(\)(.*?)do\(\)`)
	input = re.ReplaceAllString(input, "XXX")
	return input
}

func runMultInstructions(input string) int {
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := re.FindAllStringSubmatch(input, -1)

	sum := 0
	for _, match := range matches {
		// Convert the captured groups to integers
		num1 := toInt(match[1])
		num2 := toInt(match[2])
		sum += num1 * num2
	}

	return sum
}
