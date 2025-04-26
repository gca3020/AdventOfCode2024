package main

import (
	"slices"
	"strconv"
	"strings"
)

func init() {
	RegisterSolver("day1", &Day1{})
}

type Day1 struct {
	lines  []string
	l1, l2 []int
	hist   map[int]int
}

func (d *Day1) Init(input string) {
	d.lines = strings.Split(input, "\n")
	d.l1, d.l2 = d.buildLists()
	d.hist = make(map[int]int)
}

func (d *Day1) Part1() int {
	slices.Sort(d.l1)
	slices.Sort(d.l2)

	// Find the first number in l1 that is greater than the last number in l2
	sum := 0
	for i, val := range d.l1 {
		diff := d.l2[i] - val
		if diff < 0 {
			diff = -diff
		}
		sum += diff
	}
	return sum
}

func (d *Day1) buildLists() ([]int, []int) {
	l1 := []int{}
	l2 := []int{}
	for _, line := range d.lines {
		if line == "" {
			continue
		}
		tokens := strings.Fields(line)
		t1, _ := strconv.Atoi(tokens[0])
		t2, _ := strconv.Atoi(tokens[1])
		l1 = append(l1, t1)
		l2 = append(l2, t2)
	}
	return l1, l2
}

func (d *Day1) Part2() int {
	// build histogram
	for _, val := range d.l2 {
		d.hist[val]++
	}

	sum := 0
	for _, val := range d.l1 {
		if count, ok := d.hist[val]; ok {
			sum += count * val
		}
	}
	return sum
}
