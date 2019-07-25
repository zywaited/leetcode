package _316

import (
	"testing"

	"github.com/stretchr/testify/require"

	"leetcode/301_350/316/one"
)

func TestRemoveDuplicateLetters(t *testing.T) {
	require.Equal(t, "abc", one.RemoveDuplicateLetters("bcabc"))
	require.Equal(t, "acdb", one.RemoveDuplicateLetters("cbacdcbc"))
	require.Equal(t, "bac", one.RemoveDuplicateLetters("bbcaac"))
}
