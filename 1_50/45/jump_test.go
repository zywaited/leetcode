package _45

import (
	"testing"

	"github.com/stretchr/testify/require"

	"leetcode/51_100/45/ignore"
	"leetcode/51_100/45/one"
	"leetcode/51_100/45/two"
)

func TestJump(t *testing.T) {
	require.Equal(t, 2, ignore.Jump([]int{2, 3, 1, 1, 4}))
	require.Equal(t, 2, ignore.Jump([]int{1, 2, 3}))
	require.Equal(t, 2, one.Jump([]int{2, 3, 1, 1, 4}))
	require.Equal(t, 2, one.Jump([]int{1, 2, 3}))
	require.Equal(t, 2, two.Jump([]int{2, 3, 1, 1, 4}))
	require.Equal(t, 2, two.Jump([]int{1, 2, 3}))
}
