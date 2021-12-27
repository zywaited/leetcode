package _10

import (
	"leetcode/1_500/1_50/10/one"
	"leetcode/1_500/1_50/10/two"
	"testing"

	"github.com/stretchr/testify/require"
)

type isMatch func(string, string) bool

func test(t *testing.T, fn isMatch) {
	require.NotEmpty(t, fn("aa", "a*"))
	require.NotEmpty(t, fn("aa", "a."))
	require.Empty(t, fn("acdcb", "a*c.b"))
	require.NotEmpty(t, fn("bbbaccbbbaababbac", ".b*b*.*...*.*c*."))
	require.NotEmpty(t, fn("abcd", ".*ab*be*cd"))
}

func TestIsMatch(t *testing.T) {
	test(t, one.IsMatch)
	test(t, two.IsMatch)
}
