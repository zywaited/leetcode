package _55

import (
	"testing"

	"github.com/stretchr/testify/require"

	"leetcode/51_100/55/ignore"
	"leetcode/51_100/55/ignore_two"
	"leetcode/51_100/55/one"
	"leetcode/51_100/55/two"
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
