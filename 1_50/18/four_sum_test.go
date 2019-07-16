package _18

import (
	"testing"

	"leetcode/1_50/18/one"

	"github.com/stretchr/testify/require"
)

func TestFourSum(t *testing.T) {
	outs := [][]int{
		{-1, 0, 0, 1},
		{-2, -1, 1, 2},
		{-2, 0, 0, 2},
	}
	finds := one.FourSum([]int{1, 0, -1, 0, -2, 2}, 0)
	require.Len(t, finds, len(outs))
	for _, out := range outs {
		require.Contains(t, finds, out)
	}
}
