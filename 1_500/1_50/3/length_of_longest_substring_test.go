package _3

import (
	"leetcode/1_500/1_50/3/one"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	require.Equal(t, 3, one.LengthOfLongestSubstring("abcabcbb"))
	require.Equal(t, 1, one.LengthOfLongestSubstring("bbbbb"))
}
