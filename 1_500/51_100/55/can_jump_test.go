package _55

import (
	"leetcode/1_500/51_100/55/ignore"
	"leetcode/1_500/51_100/55/ignore_two"
	"leetcode/1_500/51_100/55/one"
	"leetcode/1_500/51_100/55/two"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCanJump(t *testing.T) {
	require.NotEmpty(t, ignore.CanJump([]int{2, 3, 1, 1, 4}))
	require.Empty(t, ignore.CanJump([]int{3, 2, 1, 0, 4}))

	require.NotEmpty(t, one.CanJump([]int{2, 3, 1, 1, 4}))
	require.Empty(t, one.CanJump([]int{3, 2, 1, 0, 4}))

	require.NotEmpty(t, two.CanJump([]int{2, 3, 1, 1, 4}))
	require.Empty(t, two.CanJump([]int{3, 2, 1, 0, 4}))

	require.NotEmpty(t, ignore_two.CanJump([]int{2, 3, 1, 1, 4}))
	require.Empty(t, ignore_two.CanJump([]int{3, 2, 1, 0, 4}))
	require.Empty(t, ignore_two.CanJump([]int{0, 1, 2}))
	require.NotEmpty(t, ignore_two.CanJump([]int{0}))
}
