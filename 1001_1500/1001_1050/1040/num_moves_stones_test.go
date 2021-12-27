package _1040

import (
	"leetcode/1001_1500/1001_1050/1040/one"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNumMovesStones(t *testing.T) {
	require.Equal(t, []int{1, 2}, one.NumMovesStonesII([]int{7, 4, 9}))
	require.Equal(t, []int{2, 3}, one.NumMovesStonesII([]int{6, 5, 4, 3, 10}))
	require.Equal(t, []int{0, 0}, one.NumMovesStonesII([]int{100, 101, 104, 102, 103}))
}
