package _316

import (
	"leetcode/1_500/301_350/316/one"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRemoveDuplicateLetters(t *testing.T) {
	require.Equal(t, "abc", one.RemoveDuplicateLetters("bcabc"))
	require.Equal(t, "acdb", one.RemoveDuplicateLetters("cbacdcbc"))
	require.Equal(t, "bac", one.RemoveDuplicateLetters("bbcaac"))
}
