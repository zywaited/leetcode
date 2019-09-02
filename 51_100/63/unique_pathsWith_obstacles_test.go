package _63

import (
	"leetcode/51_100/63/one"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUniquePathsWithObstacles(t *testing.T) {
	require.Equal(
		t,
		7,
		one.UniquePathsWithObstacles(
			[][]int{
				{0, 0, 0, 0},
				{0, 1, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 1, 0},
				{0, 0, 0, 0},
			},
		),
	)
}
