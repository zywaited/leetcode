package _41

import (
	"leetcode/1_500/1_50/41/one"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFirstMissingPositive(t *testing.T) {
	require.Equal(t, 3, one.FirstMissingPositive([]int{1, 2, 0}))
	require.Equal(t, 2, one.FirstMissingPositive([]int{3, 4, -1, 1}))
}
