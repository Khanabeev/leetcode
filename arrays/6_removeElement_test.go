package arrays

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name   string
		get    []int
		remove int
		want   int
	}{
		{"test-1", []int{3, 2, 2, 3}, 2, 2},
		{"test-1", []int{0, 0, 1, 1}, 0, 2},
		{"test-1", []int{0, 0, 1, 0}, 0, 1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.want, removeElement(test.get, test.remove))
		})
	}
}
