package _1345

import (
	"leetcode/1001_1500/1301_1350/1345/one"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMinJumps(t *testing.T) {
	require.Equal(t, one.MinJumps([]int{100, -23, -23, 404, 100, 23, 23, 23, 3, 404}), 3)
	require.Equal(t, one.MinJumps([]int{7}), 0)
	require.Equal(t, one.MinJumps([]int{7, 6, 9, 6, 9, 6, 9, 7}), 1)
	require.Equal(t, one.MinJumps([]int{6, 1, 9}), 2)
	require.Equal(t, one.MinJumps([]int{11, 22, 7, 7, 7, 7, 7, 7, 7, 22, 13}), 3)
	require.Equal(t, one.MinJumps([]int{0, 4, 3, 9}), 3)
}
