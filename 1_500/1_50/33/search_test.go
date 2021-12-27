package _33

import (
	"leetcode/1_500/1_50/33/one"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSearch(t *testing.T) {
	require.Equal(t, 4, one.Search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	require.Equal(t, -1, one.Search([]int{4, 5, 6, 7, 0, 1, 2}, 3))
}
