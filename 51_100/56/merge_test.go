package _56

import (
	"testing"

	"leetcode/51_100/56/one"
	"leetcode/51_100/56/two"

	"github.com/stretchr/testify/require"
)

type merge func([][]int) [][]int

func test(t *testing.T, fn merge) {
	outs := [][]int{
		{1, 6},
		{8, 10},
		{15, 18},
	}
	mgs := fn([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}})
	require.Len(t, mgs, len(outs))
	for _, out := range outs {
		require.Contains(t, mgs, out)
	}

	require.Equal(t, [][]int{{1, 10}}, fn([][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}}))
	require.Equal(t, [][]int{{0, 0}, {1, 4}}, fn([][]int{{1, 4}, {0, 0}}))
	require.Equal(t, [][]int{{1, 3}, {4, 7}}, fn([][]int{{2, 3}, {2, 2}, {3, 3}, {1, 3}, {5, 7}, {2, 2}, {4, 6}}))
	require.Equal(t, [][]int{{2, 7}}, fn([][]int{{2, 3}, {4, 6}, {5, 7}, {3, 4}}))
}

func TestMerge(t *testing.T) {
	test(t, one.Merge)
	test(t, two.Merge)
}
