package _4

import (
	"leetcode/1_500/1_50/4/one"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindMedianSortedArrays(t *testing.T) {
	require.Equal(t, float64(2.0), one.FindMedianSortedArrays([]int{1, 3}, []int{2}))
	require.Equal(t, float64(2.5), one.FindMedianSortedArrays([]int{1, 3}, []int{2, 4}))
}
