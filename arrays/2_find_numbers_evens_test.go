package arrays

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindNumbers(t *testing.T) {
	tests := []struct {
		name string
		get  []int
		want int
	}{
		{"Check 3 evens", []int{23, 24, 45}, 3},
		{"Check 1 evens", []int{23, 243, 451}, 1},
		{"Check 0 evens", []int{223, 243, 451}, 0},
		{"Check 1 even", []int{2231}, 1},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.want, findNumbers(test.get))
		})
	}
}
