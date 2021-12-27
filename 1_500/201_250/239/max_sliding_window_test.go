package one

import (
	"leetcode/1_500/201_250/239/one"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaxSlidingWindow(t *testing.T) {
	require.Equal(t, []int{3, 3, 5, 5, 6, 7}, one.MaxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}
