package main

import (
	"fmt"
	"slices"
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	for _, size := range []int{-10, -1, 0, 1, 7, 8, 9, 1_000} {
		t.Run(fmt.Sprintf("size_%d", size), func(t *testing.T) {
			data := generateRandomElements(size)
			wantLength := size
			if size < 0 {
				wantLength = 0
			}
			if len(data) != wantLength {
				t.Fatalf("length = %d, want %d", len(data), wantLength)
			}
			for i, value := range data {
				if value <= 0 {
					t.Errorf("data[%d] = %d, want a positive integer", i, value)
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	testMaximum(t, maximum)
}

func TestMaxChunks(t *testing.T) {
	testMaximum(t, maxChunks)
}

func testMaximum(t *testing.T, findMaximum func([]int) int) {
	t.Helper()
	maxInt := int(^uint(0) >> 1)
	minInt := -maxInt - 1
	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "nil", data: nil, want: 0},
		{name: "empty", data: []int{}, want: 0},
		{name: "single", data: []int{42}, want: 42},
		{name: "single_negative", data: []int{-42}, want: -42},
		{name: "zeros", data: []int{0, 0, 0}, want: 0},
		{name: "equal", data: []int{5, 5, 5}, want: 5},
		{name: "ascending", data: []int{1, 2, 3, 4, 5}, want: 5},
		{name: "descending", data: []int{5, 4, 3, 2, 1}, want: 5},
		{name: "middle", data: []int{2, 9, 3}, want: 9},
		{name: "negative", data: []int{-9, -2, -7}, want: -2},
		{name: "mixed", data: []int{-9, 0, 7, -3}, want: 7},
		{name: "integer_limits", data: []int{minInt, 0, maxInt}, want: maxInt},
		{name: "minimum_integer", data: []int{minInt}, want: minInt},
		{name: "negative_chunks", data: []int{-9, -8, -7, -6, -5, -4, -3, -2, -1}, want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			original := slices.Clone(tt.data)
			if got := findMaximum(tt.data); got != tt.want {
				t.Errorf("maximum = %d, want %d", got, tt.want)
			}
			if !slices.Equal(tt.data, original) {
				t.Error("input slice was modified")
			}
		})
	}
}

func TestMaxChunksEveryPosition(t *testing.T) {
	// Exercise short slices, chunk boundaries and every possible remainder.
	for size := 1; size <= 4*CHUNKS+1; size++ {
		for position := 0; position < size; position++ {
			data := make([]int, size)
			for i := range data {
				data[i] = -10
			}
			data[position] = -1
			if got := maxChunks(data); got != -1 {
				t.Fatalf("size %d, maximum at %d: got %d, want -1", size, position, got)
			}
		}
	}
}
