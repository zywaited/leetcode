package _84

import (
	"leetcode/1_500/51_100/84/one"
	"leetcode/1_500/51_100/84/two"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLargestRectangleArea(t *testing.T) {
	require.Equal(t, one.LargestRectangleArea([]int{2, 1, 5, 6, 2, 3}), two.LargestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
}
