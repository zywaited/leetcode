package _1187

import (
	"leetcode/1001_1500/1151_1200/1187/one"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMakeArrayIncreasing(t *testing.T) {
	require.Equal(t, 1, one.MakeArrayIncreasing([]int{1, 5, 3, 6, 7}, []int{1, 3, 2, 4}))
	require.Equal(t, 0, one.MakeArrayIncreasing([]int{1, math.MaxInt32}, []int{2}))
	require.Equal(t, -1, one.MakeArrayIncreasing([]int{1, 5, 3, 6, 7}, []int{1, 6, 3, 3}))
}
