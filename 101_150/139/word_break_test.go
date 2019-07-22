package _139

import (
	"testing"

	"github.com/stretchr/testify/require"

	"leetcode/101_150/139/one"
)

func TestWordBreak(t *testing.T) {
	require.NotEmpty(t, one.WordBreak("leetcode", []string{"leet", "code"}))
	require.NotEmpty(t, one.WordBreak("applepenapple", []string{"apple", "pen"}))
	require.Empty(t, one.WordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
	require.Empty(t, one.WordBreak("a", []string{"b"}))
}
