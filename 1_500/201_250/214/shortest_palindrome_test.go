package _214

import (
	"leetcode/1_500/201_250/214/one"
	"leetcode/1_500/201_250/214/two"
	"testing"

	"github.com/stretchr/testify/require"
)

type shortestPalindrome func(string) string

func test(t *testing.T, fn shortestPalindrome) {
	require.Equal(t, "aaacecaaa", fn("aacecaaa"))
	require.Equal(t, "dcbabcd", fn("abcd"))
}

func TestShortestPalindrome(t *testing.T) {
	test(t, one.ShortestPalindrome)
	test(t, two.ShortestPalindrome)
}
