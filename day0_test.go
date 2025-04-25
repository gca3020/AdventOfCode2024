package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input0 = ""

func TestDay0_Part1(t *testing.T) {
	d := &Day0{}
	d.Init(input0)
	assert.Equal(t, 0, d.Part1())
}

func TestDay0_Part2(t *testing.T) {
	d := &Day0{}
	d.Init(input0)
	assert.Equal(t, 0, d.Part2())
}
