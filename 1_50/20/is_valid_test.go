package _20

import (
	"testing"

	"leetcode/1_50/20/one"

	"github.com/stretchr/testify/require"
)

func TestIsValid(t *testing.T) {
	require.NotEmpty(t, one.IsValid(""))
	require.NotEmpty(t, one.IsValid("()"))
	require.NotEmpty(t, one.IsValid("()[]{}"))
	require.Empty(t, one.IsValid("(]"))
	require.Empty(t, one.IsValid("([)]"))
	require.NotEmpty(t, one.IsValid("{[]}"))
	require.Empty(t, one.IsValid("["))
}
