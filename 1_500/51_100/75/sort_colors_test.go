package _75

import (
	"leetcode/1_500/51_100/75/one"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSortColors(t *testing.T) {
	in := []int{2, 0, 2, 1, 1, 0}
	one.SortColors(in)
	require.Equal(t, []int{0, 0, 1, 1, 2, 2}, in)

	in = []int{}
	one.SortColors(in)
	require.Equal(t, []int{}, in)

	in = []int{1}
	one.SortColors(in)
	require.Equal(t, []int{1}, in)

	in = []int{1, 2, 0, 0, 1, 2, 0, 1, 1, 2, 2, 0, 0}
	one.SortColors(in)
	require.Equal(t, []int{0, 0, 0, 0, 0, 1, 1, 1, 1, 2, 2, 2, 2}, in)
}
