package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input1 = `3   4
4   3
2   5
1   3
3   9
3   3
`

func TestDay1_Part1(t *testing.T) {
	s := &Day1{}
	s.Init(input1)
	assert.Equal(t, 11, s.Part1())
}

func TestDay1_Part2(t *testing.T) {
	s := &Day1{}
	s.Init(input1)
	assert.Equal(t, 31, s.Part2())
}
