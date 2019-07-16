package _51

import (
	"testing"

	"leetcode/51_100/51/one"

	"github.com/stretchr/testify/require"
)

func TestSolveNQueens(t *testing.T) {
	answers := one.SolveNQueens(4)
	outs := [][]string{
		{
			".Q..",
			"...Q",
			"Q...",
			"..Q.",
		},
		{
			"..Q.",
			"Q...",
			"...Q",
			".Q..",
		},
	}
	require.Len(t, answers, len(outs))
	for _, out := range outs {
		require.Contains(t, answers, out)
	}
}
