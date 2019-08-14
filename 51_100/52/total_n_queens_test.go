package _52

import (
	"testing"

	"github.com/stretchr/testify/require"

	"leetcode/51_100/52/one"
)

func TestTotalNQueens(t *testing.T) {
	require.Equal(t, 2, one.TotalNQueens(4))
}
