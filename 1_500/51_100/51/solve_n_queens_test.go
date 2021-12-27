package _51

import (
	"leetcode/1_500/51_100/51/one"
	"testing"

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
