package _1040

import (
	"testing"

	"leetcode/1001_1050/1040/one"

	"github.com/stretchr/testify/require"
)

func TestNumMovesStones(t *testing.T) {
	require.Equal(t, []int{1, 2}, one.NumMovesStonesII([]int{7, 4, 9}))
	require.Equal(t, []int{2, 3}, one.NumMovesStonesII([]int{6, 5, 4, 3, 10}))
}
