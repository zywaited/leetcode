package _44

import (
	"testing"

	"leetcode/1_50/44/one"
	"leetcode/1_50/44/two"

	"github.com/stretchr/testify/require"
)

type isMatch func(string, string) bool

func test(t *testing.T, fn isMatch) {
	require.NotEmpty(t, fn("aa", "a*"))
	require.NotEmpty(t, fn("aa", "a?"))
	require.NotEmpty(t, fn("adceb", "*a*b"))
	require.Empty(t, fn("acdcb", "a*c?b"))
}

func TestIsMatch(t *testing.T) {
	test(t, one.IsMatch)
	test(t, two.IsMatch)
}
