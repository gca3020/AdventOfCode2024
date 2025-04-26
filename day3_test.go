package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input3 = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
var input3_p2 = "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"

func TestDay3_Part1(t *testing.T) {
	d := &Day3{}
	d.Init(input3)
	assert.Equal(t, 161, d.Part1())
}

func TestDay3_Part2(t *testing.T) {
	d := &Day3{}
	d.Init(input3_p2)
	assert.Equal(t, 48, d.Part2())
}

func TestTrimInstructions(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "No match",
			input:    "mul(2,4)mul(3,7)",
			expected: "mul(2,4)mul(3,7)",
		},
		{
			name:     "Example",
			input:    "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
			expected: "xmul(2,4)&mul[3,7]!^XXX?mul(8,5))",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := trimInstructions(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
