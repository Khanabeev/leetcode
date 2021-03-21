package arrays

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		name string
		get1 []int
		len1 int
		get2 []int
		len2 int
		want []int
	}{
		{"test-1", []int{1, 2, 3, 0, 0, 0}, 3, []int{4, 5, 6}, 3, []int{1, 2, 3, 4, 5, 6}},
		{"test-2", []int{1, 0, 0, 0}, 1, []int{4}, 1, []int{1, 4, 0, 0}},
		{"test-3", []int{}, 0, []int{}, 0, []int{}},
		{"test-4", []int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			arr := merge(test.get1, test.len1, test.get2, test.len2)
			require.Equal(t, test.want, arr)
		})
	}
}
