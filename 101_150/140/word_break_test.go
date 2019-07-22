package _140

import (
	"testing"

	"github.com/stretchr/testify/require"

	"leetcode/101_150/140/one"
)

func test(t *testing.T, outs []string, finds []string) {
	require.Len(t, finds, len(outs))
	for _, out := range outs {
		require.Contains(t, finds, out)
	}
}

func TestWordBreak(t *testing.T) {
	test(
		t,
		[]string{
			"cats and dog",
			"cat sand dog",
		},
		one.WordBreak("catsanddog", []string{"cat", "cats", "and", "sand", "dog"}),
	)

	test(
		t,
		[]string{
			"pine apple pen apple",
			"pineapple pen apple",
			"pine applepen apple",
		},
		one.WordBreak("pineapplepenapple", []string{"apple", "pen", "applepen", "pine", "pineapple"}),
	)

	test(
		t,
		nil,
		one.WordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}),
	)
}
