package _1210

import (
	"testing"

	"github.com/stretchr/testify/require"

	"leetcode/1201_1250/1210/one"
	"leetcode/1201_1250/1210/two"
)

func TestMinimunMoves(t *testing.T) {
	require.Equal(
		t,
		11,
		one.MinimumMoves([][]int{
			{0, 0, 0, 0, 0, 1},
			{1, 1, 0, 0, 1, 0},
			{0, 0, 0, 0, 1, 1},
			{0, 0, 1, 0, 1, 0},
			{0, 1, 1, 0, 0, 0},
			{0, 1, 1, 0, 0, 0},
		}),
	)
	require.Equal(
		t,
		9,
		one.MinimumMoves([][]int{
			{0, 0, 1, 1, 1, 1},
			{0, 0, 0, 0, 1, 1},
			{1, 1, 0, 0, 0, 1},
			{1, 1, 1, 0, 0, 1},
			{1, 1, 1, 0, 0, 1},
			{1, 1, 1, 0, 0, 0},
		}),
	)

	require.Equal(
		t,
		11,
		two.MinimumMoves([][]int{
			{0, 0, 0, 0, 0, 1},
			{1, 1, 0, 0, 1, 0},
			{0, 0, 0, 0, 1, 1},
			{0, 0, 1, 0, 1, 0},
			{0, 1, 1, 0, 0, 0},
			{0, 1, 1, 0, 0, 0},
		}),
	)
	require.Equal(
		t,
		9,
		two.MinimumMoves([][]int{
			{0, 0, 1, 1, 1, 1},
			{0, 0, 0, 0, 1, 1},
			{1, 1, 0, 0, 0, 1},
			{1, 1, 1, 0, 0, 1},
			{1, 1, 1, 0, 0, 1},
			{1, 1, 1, 0, 0, 0},
		}),
	)
}
