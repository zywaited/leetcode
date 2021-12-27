package _1354

import (
	"leetcode/1001_1500/1351_1400/1354/one"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsPossible(t *testing.T) {
	require.Empty(t, one.IsPossible([]int{9, 2, 5}))
	require.NotEmpty(t, one.IsPossible([]int{9, 3, 5}))
	require.NotEmpty(t, one.IsPossible([]int{100000000, 1}))
	require.NotEmpty(t, one.IsPossible([]int{8, 5}))
	require.Empty(t, one.IsPossible([]int{1, 1, 2}))
}
