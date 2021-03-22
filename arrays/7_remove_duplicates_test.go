package arrays

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name string
		get  []int
		want int
	}{
		{"test-1", []int{3, 2, 2, 3}, 2},
		{"test-2", []int{1, 1, 2}, 2},
		{"test-3", []int{1, 2, 3, 4}, 4},
		{"test-4", []int{}, 0},
		{"test-5", []int{1, 1, 1}, 1},
		{"test-6", []int{0, 0, 0, 0, 0}, 1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.want, removeDuplicates(test.get))
		})
	}
}
