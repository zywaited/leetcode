package _46

import (
	"leetcode/1_500/1_50/46/one"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPermute(t *testing.T) {
	outs := [][]int{
		{1, 2, 3},
		{1, 3, 2},
		{2, 1, 3},
		{2, 3, 1},
		{3, 1, 2},
		{3, 2, 1},
	}
	finds := one.Permute([]int{1, 2, 3})
	require.Len(t, finds, len(outs))
	for _, out := range outs {
		require.Contains(t, finds, out)
	}
}
