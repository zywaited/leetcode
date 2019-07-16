package _32

import (
	"testing"

	"leetcode/1_50/32/one"

	"github.com/stretchr/testify/require"
)

func TestLongestValidParentheses(t *testing.T) {
	require.Equal(t, 2, one.LongestValidParentheses("(()"))
	require.Equal(t, 4, one.LongestValidParentheses(")()())"))
}
