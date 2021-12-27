package _34

import (
	"leetcode/1_500/1_50/34/one"
	"leetcode/1_500/1_50/34/two"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSearchRange(t *testing.T) {
	require.Equal(t, []int{3, 4}, one.SearchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	require.Equal(t, []int{-1, -1}, one.SearchRange([]int{5, 7, 7, 8, 8, 10}, 6))
	require.Equal(t, []int{1, 2}, one.SearchRange([]int{5, 7, 7, 8, 8, 10}, 7))
	require.Equal(t, []int{0, 0}, one.SearchRange([]int{5}, 5))
	require.Equal(t, []int{-1, -1}, one.SearchRange([]int{5}, 6))
	require.Equal(t, []int{1, 4}, one.SearchRange([]int{5, 7, 7, 7, 7, 10}, 7))
	require.Equal(t, []int{0, 5}, one.SearchRange([]int{7, 7, 7, 7, 7, 7}, 7))

	require.Equal(t, []int{3, 4}, two.SearchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	require.Equal(t, []int{-1, -1}, two.SearchRange([]int{5, 7, 7, 8, 8, 10}, 6))
	require.Equal(t, []int{1, 2}, two.SearchRange([]int{5, 7, 7, 8, 8, 10}, 7))
	require.Equal(t, []int{0, 0}, two.SearchRange([]int{5}, 5))
	require.Equal(t, []int{-1, -1}, two.SearchRange([]int{5}, 6))
	require.Equal(t, []int{1, 4}, two.SearchRange([]int{5, 7, 7, 7, 7, 10}, 7))
	require.Equal(t, []int{0, 5}, two.SearchRange([]int{7, 7, 7, 7, 7, 7}, 7))
	require.Equal(t, []int{-1, -1}, two.SearchRange([]int{}, 0))
}
