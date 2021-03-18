package arrays

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFindMaxConsecutiveOnes(t *testing.T) {
	tests := []struct {
		name string
		get  []int
		want int
	}{
		{
			name: "Check 10 1s",
			get:  []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			want: 10,
		},
		{
			name:
			"Check 2 1s",
			get:  []int{1, 1, 0, 1, 0, 1, 1, 0, 1, 1},
			want: 2,
		},
		{
			name:
			"Check 0 1s",
			get:  []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			want: 0,
		},
		{
			name:
			"Check 1 1s",
			get:  []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
			want: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.want, findMaxConsecutiveOnes(test.get))
		})
	}
}
