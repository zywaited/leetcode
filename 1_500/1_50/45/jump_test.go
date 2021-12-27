package _45

import (
	"leetcode/1_500/1_50/45/ignore"
	"leetcode/1_500/1_50/45/one"
	"leetcode/1_500/1_50/45/two"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestJump(t *testing.T) {
	require.Equal(t, 2, ignore.Jump([]int{2, 3, 1, 1, 4}))
	require.Equal(t, 2, ignore.Jump([]int{1, 2, 3}))
	require.Equal(t, 2, one.Jump([]int{2, 3, 1, 1, 4}))
	require.Equal(t, 2, one.Jump([]int{1, 2, 3}))
	require.Equal(t, 2, two.Jump([]int{2, 3, 1, 1, 4}))
	require.Equal(t, 2, two.Jump([]int{1, 2, 3}))
}
