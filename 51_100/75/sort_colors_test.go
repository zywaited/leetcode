package _75

import (
	"testing"

	"github.com/stretchr/testify/require"

	"leetcode/51_100/75/one"
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
