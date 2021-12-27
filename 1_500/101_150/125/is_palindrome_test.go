package _125

import (
	"leetcode/1_500/101_150/125/one"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsPalindrome(t *testing.T) {
	require.NotEmpty(t, one.IsPalindrome("A man, a plan, a canal: Panama"))
	require.NotEmpty(t, one.IsPalindrome(""))
	require.NotEmpty(t, one.IsPalindrome("a"))
	require.Empty(t, one.IsPalindrome("race a car"))
	require.NotEmpty(t, one.IsPalindrome("A man, a plan, a c1c anal: Panama"))
}
