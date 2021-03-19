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
		{"test-3", []int{1,0,2,3,0,4,5,0}, []int{1,0,0,2,3,0,0,4}},
		{"test-4", []int{1,0,0,3,0,4,5,0}, []int{1,0,0,0,0,3,0,0}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.want, duplicateZeros(test.get))
		})
	}
}
