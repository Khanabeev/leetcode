package arrays

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDuplicateZeros(t *testing.T) {
	tests := []struct {
		name string
		get  []int
		want []int
	}{
		{"test-1", []int{4, 2, 3, 6}, []int{4, 2, 3, 6}},
		{"test-2", []int{2, 0, 4, 0, 1}, []int{2, 0, 0, 4, 0}},
		{"test-3", []int{1, 0, 2, 3, 0, 4, 5, 0}, []int{1, 0, 0, 2, 3, 0, 0, 4}},
		{"test-4", []int{1, 0, 3, 0, 4, 5, 0}, []int{1, 0, 0, 3, 0, 0, 4}},
		{"test-5", []int{0, 0, 3, 1, 2, 3, 4}, []int{0, 0, 0, 0, 3, 1, 2}},
		{"test-6", []int{0, 0, 0, 0, 2, 3, 4}, []int{0, 0, 0, 0, 0, 0, 0}},
		{"test-7", []int{}, []int{}},
		{"test-8", []int{8,4,5,0,0,0,0,7}, []int{8,4,5,0,0,0,0,0}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.ElementsMatch(t, test.want, duplicateZeros(test.get))
		})
	}
}
