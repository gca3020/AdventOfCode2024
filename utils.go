package main

import (
	"log/slog"
	"strconv"
	"strings"
)

func toInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		slog.Error("Could not convert to integer", "str", str, "err", err)
		panic("Error converting string to integer")
	}
	return i
}

func toIntSlice(s string) []int {
	strs := strings.Fields(s)
	ints := make([]int, 0, len(strs))
	for _, s := range strs {
		ints = append(ints, toInt(s))
	}
	return ints
}

func absDiff(a, b int) int {
	d := a - b
	if d < 0 {
		d = -d
	}
	return d
}

func between(val, a, b int) bool {
	return (val > a && val < b) || (val < a && val > b)
}
