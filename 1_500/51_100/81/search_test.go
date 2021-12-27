package _81

import (
	"leetcode/1_500/51_100/81/one"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSearch(t *testing.T) {
	require.NotEmpty(t, one.Search([]int{2, 5, 6, 0, 0, 1, 2}, 0))
	require.Empty(t, one.Search([]int{2, 5, 6, 0, 0, 1, 2}, 3))
	require.NotEmpty(t, one.Search([]int{1, 1, 2, 3, 1}, 3))
	require.NotEmpty(t, one.Search([]int{1, 1, 1, 3, 1}, 3))
	require.NotEmpty(t, one.Search([]int{1, 3, 1, 1, 1}, 3))
}
