package _139

import (
	"leetcode/1_500/101_150/139/one"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWordBreak(t *testing.T) {
	require.NotEmpty(t, one.WordBreak("leetcode", []string{"leet", "code"}))
	require.NotEmpty(t, one.WordBreak("applepenapple", []string{"apple", "pen"}))
	require.Empty(t, one.WordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
	require.Empty(t, one.WordBreak("a", []string{"b"}))
}
