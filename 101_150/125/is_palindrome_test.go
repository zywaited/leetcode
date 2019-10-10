package _125

import (
	"testing"

	"github.com/stretchr/testify/require"

	"leetcode/101_150/125/one"
)

func TestIsPalindrome(t *testing.T) {
	require.NotEmpty(t, one.IsPalindrome("A man, a plan, a canal: Panama"))
	require.NotEmpty(t, one.IsPalindrome(""))
	require.NotEmpty(t, one.IsPalindrome("a"))
	require.Empty(t, one.IsPalindrome("race a car"))
	require.NotEmpty(t, one.IsPalindrome("A man, a plan, a c1c anal: Panama"))
}
