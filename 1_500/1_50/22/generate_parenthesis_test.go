package _22

import (
	"leetcode/1_500/1_50/22/one"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateParenthesis(t *testing.T) {
	ex := []string{
		"((()))",
		"(()())",
		"(())()",
		"()(())",
		"()()()",
	}
	ac := one.GenerateParenthesis(3)
	require.Len(t, ac, len(ex))
	for _, e := range ex {
		require.Contains(t, ac, e)
	}
}
