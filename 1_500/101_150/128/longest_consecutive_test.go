package _128

import (
	"leetcode/1_500/101_150/128/one"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLongestConsecutive(t *testing.T) {
	require.Equal(t, 4, one.LongestConsecutive([]int{100, 4, 200, 1, 3, 2}))
}
