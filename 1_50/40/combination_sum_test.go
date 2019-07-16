package _40

import (
	"testing"

	"leetcode/1_50/40/one"

	"github.com/stretchr/testify/require"
)

func TestCombinationSum(t *testing.T) {
	finds := one.CombinationSum([]int{10, 1, 2, 7, 6, 1, 5}, 8)
	outs := [][]int{
		{1, 7},
		{1, 2, 5},
		{2, 6},
		{1, 1, 6},
	}
	require.Len(t, finds, len(outs))
	for _, out := range outs {
		require.Contains(t, finds, out)
	}
}
