package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input2 = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9
`

func TestDay2_Part1(t *testing.T) {
	s := &Day2{}
	s.Init(input2)
	assert.Equal(t, 2, s.Part1())
}

func TestDay2_Part2(t *testing.T) {
	s := &Day2{}
	s.Init(input2)
	assert.Equal(t, 4, s.Part2())
}

func TestReport_Safe(t *testing.T) {
	tests := []struct {
		name     string
		levels   []int
		expected bool
	}{
		{
			name:     "Safe sequence",
			levels:   []int{1, 3, 6, 7, 9},
			expected: true,
		},
		{
			name:     "Safe reverse sequence",
			levels:   []int{7, 6, 4, 2, 1},
			expected: true,
		},
		{
			name:     "Unsafe due to large difference",
			levels:   []int{1, 2, 6, 7, 8},
			expected: false,
		},
		{
			name:     "Unsafe due to not between",
			levels:   []int{1, 3, 2, 4, 6},
			expected: false,
		},
		{
			name:     "Unsafe due to too large gap",
			levels:   []int{45, 48, 51, 52, 55, 58, 62},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			report := Report{levels: tt.levels}
			assert.Equal(t, tt.expected, report.Safe())
		})
	}
}

func TestNewRemovingIndex(t *testing.T) {
	tests := []struct {
		name     string
		levels   []int
		index    int
		expected []int
	}{
		{
			name:     "Remove middle element",
			levels:   []int{1, 2, 3, 4, 5},
			index:    2,
			expected: []int{1, 2, 4, 5},
		},
		{
			name:     "Remove first element",
			levels:   []int{1, 2, 3, 4, 5},
			index:    0,
			expected: []int{2, 3, 4, 5},
		},
		{
			name:     "Remove last element",
			levels:   []int{1, 2, 3, 4, 5},
			index:    4,
			expected: []int{1, 2, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			report := Report{levels: tt.levels}
			result := NewRemovingIndex(report, tt.index)
			assert.Equal(t, tt.expected, result.levels)
		})
	}
}
