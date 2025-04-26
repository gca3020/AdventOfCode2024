package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_toIntSlice(t *testing.T) {
	tests := []struct {
		name      string
		args      string
		want      []int
		wantPanic bool
	}{
		{
			name: "Single number",
			args: "42",
			want: []int{42},
		},
		{
			name: "Multiple numbers",
			args: "1 2 3 4 5",
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "Numbers with extra spaces",
			args: "  10   20  30 ",
			want: []int{10, 20, 30},
		},
		{
			name: "Empty string",
			args: "",
			want: []int{},
		},
		{
			name:      "Invalid number",
			args:      "1 2 three 4",
			want:      nil,
			wantPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantPanic {
				assert.Panics(t, func() { toIntSlice(tt.args) })
			} else {
				assert.ElementsMatch(t, toIntSlice(tt.args), tt.want)
			}
		})
	}
}

func Test_absDiff(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{
			name: "Positive difference",
			a:    10,
			b:    5,
			want: 5,
		},
		{
			name: "Negative difference",
			a:    5,
			b:    10,
			want: 5,
		},
		{
			name: "Zero difference",
			a:    7,
			b:    7,
			want: 0,
		},
		{
			name: "Large positive difference",
			a:    1000,
			b:    500,
			want: 500,
		},
		{
			name: "Large negative difference",
			a:    500,
			b:    1000,
			want: 500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := absDiff(tt.a, tt.b)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_between(t *testing.T) {
	tests := []struct {
		name string
		val  int
		a    int
		b    int
		want bool
	}{
		{
			name: "Value between a and b",
			val:  5,
			a:    1,
			b:    10,
			want: true,
		},
		{
			name: "Value between b and a",
			val:  5,
			a:    10,
			b:    1,
			want: true,
		},
		{
			name: "Value equal to a",
			val:  1,
			a:    1,
			b:    10,
			want: false,
		},
		{
			name: "Value equal to b",
			val:  10,
			a:    1,
			b:    10,
			want: false,
		},
		{
			name: "Value outside range",
			val:  15,
			a:    1,
			b:    10,
			want: false,
		},
		{
			name: "Value outside reversed range",
			val:  -5,
			a:    10,
			b:    1,
			want: false,
		},
		{
			name: "Value between negative range",
			val:  -5,
			a:    -10,
			b:    -1,
			want: true,
		},
		{
			name: "Value outside negative range",
			val:  -15,
			a:    -10,
			b:    -1,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := between(tt.val, tt.a, tt.b)
			assert.Equal(t, tt.want, got)
		})
	}
}
