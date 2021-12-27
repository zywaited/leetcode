package _1340

import (
	"leetcode/1001_1500/1301_1350/1340/one"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaxJumps(t *testing.T) {
	require.Equal(t, 4, one.MaxJumps([]int{6, 4, 14, 6, 8, 13, 9, 7, 10, 6, 12}, 2))
	require.Equal(t, 1, one.MaxJumps([]int{3, 3, 3, 3, 3}, 3))
	require.Equal(t, 7, one.MaxJumps([]int{7, 6, 5, 4, 3, 2, 1}, 7))
	require.Equal(t, 2, one.MaxJumps([]int{7, 1, 7, 1, 7, 1}, 2))
	require.Equal(t, 1, one.MaxJumps([]int{66}, 1))
}
