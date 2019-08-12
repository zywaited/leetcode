package _22

import (
	"testing"

	"github.com/stretchr/testify/require"

	"leetcode/1_50/22/one"
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
