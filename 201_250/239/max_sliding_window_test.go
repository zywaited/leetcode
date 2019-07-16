package one

import (
	"testing"

	"leetcode/201_250/239/one"

	"github.com/stretchr/testify/require"
)

func TestMaxSlidingWindow(t *testing.T) {
	require.Equal(t, []int{3, 3, 5, 5, 6, 7}, one.MaxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}
