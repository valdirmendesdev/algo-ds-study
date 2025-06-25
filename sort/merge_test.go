package sort_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/valdirmendesdev/algo-ds-study/sort"
)

func TestMergeSort(t *testing.T) {
	// Test cases for merge sort
	tests := []struct {
		input    []int
		expected []int
	}{
		{input: []int{5, 2, 9, 1, 5, 6}, expected: []int{1, 2, 5, 5, 6, 9}},
		{input: []int{3, 0, -2, 8, -1}, expected: []int{-2, -1, 0, 3, 8}},
		{input: []int{}, expected: []int{}},
	}

	for _, test := range tests {
		result := sort.MergeSort(test.input)
		if !assert.Equal(t, result, test.expected) {
			t.Errorf("mergeSort(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}
