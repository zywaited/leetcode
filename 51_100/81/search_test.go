package _81

import (
	"testing"

	"github.com/stretchr/testify/require"

	"leetcode/51_100/81/one"
)

func TestSearch(t *testing.T) {
	require.NotEmpty(t, one.Search([]int{2, 5, 6, 0, 0, 1, 2}, 0))
	require.Empty(t, one.Search([]int{2, 5, 6, 0, 0, 1, 2}, 3))
	require.NotEmpty(t, one.Search([]int{1, 1, 2, 3, 1}, 3))
	require.NotEmpty(t, one.Search([]int{1, 1, 1, 3, 1}, 3))
	require.NotEmpty(t, one.Search([]int{1, 3, 1, 1, 1}, 3))
}
