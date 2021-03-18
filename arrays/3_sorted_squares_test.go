package arrays

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSortedSquares(t *testing.T) {
	tests := []struct {
		name string
		get  []int
		want []int
	}{
		{"test-0", []int{-7, -3, 2, 3, 11}, []int{4, 9, 9, 49, 121}},
		{"test-1", []int{1}, []int{1}},
		{"test-2", []int{2,1}, []int{1,4}},
	}

	for _, test := range tests{
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.want, sortedSquares(test.get))
		})
	}
}
