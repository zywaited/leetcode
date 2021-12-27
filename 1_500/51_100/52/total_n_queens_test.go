package _52

import (
	"leetcode/1_500/51_100/52/one"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTotalNQueens(t *testing.T) {
	require.Equal(t, 2, one.TotalNQueens(4))
}
