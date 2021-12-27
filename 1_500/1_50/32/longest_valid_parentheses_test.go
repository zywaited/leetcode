package _32

import (
	"leetcode/1_500/1_50/32/one"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLongestValidParentheses(t *testing.T) {
	require.Equal(t, 2, one.LongestValidParentheses("(()"))
	require.Equal(t, 4, one.LongestValidParentheses(")()())"))
}
