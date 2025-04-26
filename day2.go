package main

import (
	"slices"
	"strings"
)

func init() {
	RegisterSolver("day2", &Day2{})
}

type Day2 struct {
	reports []Report
}

func (d *Day2) Init(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	d.reports = nil
	for _, line := range lines {
		d.reports = append(d.reports, NewReport(line))
	}
}

func (d *Day2) Part1() int {
	total := 0
	for _, report := range d.reports {
		if report.Safe() {
			total++
		}
	}
	return total
}

func (d *Day2) Part2() int {
	total := 0
	for _, report := range d.reports {
		if report.SafeWithDampener() {
			total++
		}
	}
	return total
}

type Report struct {
	levels []int
}

func NewReport(line string) Report {
	return Report{
		levels: toIntSlice(line),
	}
}

func NewRemovingIndex(r Report, i int) Report {
	newLevels := make([]int, len(r.levels))
	copy(newLevels, r.levels)
	return Report{
		levels: slices.Delete(newLevels, i, i+1),
	}
}

func (r Report) SafeWithDampener() bool {
	if r.Safe() {
		return true
	}

	for i := 0; i < len(r.levels); i++ {
		newReport := NewRemovingIndex(r, i)
		if newReport.Safe() {
			return true
		}
	}
	return false
}

func (r Report) Safe() bool {
	last := r.levels[len(r.levels)-1]
	for i := 1; i < len(r.levels); i++ {
		cur := r.levels[i]
		prev := r.levels[i-1]
		if absDiff(cur, prev) > 3 {
			return false
		}
		if i == len(r.levels)-1 {
			continue
		}
		if !between(cur, prev, last) {
			return false
		}
	}

	return true
}
