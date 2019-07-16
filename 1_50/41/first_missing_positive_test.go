package _41

import (
	"testing"

	"leetcode/1_50/41/one"

	"github.com/stretchr/testify/require"
)

func TestFirstMissingPositive(t *testing.T) {
	require.Equal(t, 3, one.FirstMissingPositive([]int{1, 2, 0}))
	require.Equal(t, 2, one.FirstMissingPositive([]int{3, 4, -1, 1}))
}
