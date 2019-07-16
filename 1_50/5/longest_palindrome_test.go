package _5

import (
	"testing"

	"leetcode/1_50/5/one"
	"leetcode/1_50/5/two"

	"github.com/stretchr/testify/require"
)

type longestPalindrome func(string) string

func test(t *testing.T, fn longestPalindrome) {
	require.Contains(t, []string{"bab", "aba"}, fn("babad"))
	require.Equal(t, "bb", fn("cbbd"))
}

func TestLongestPalindrome(t *testing.T) {
	test(t, one.LongestPalindrome)
	test(t, two.LongestPalindrome)
}
