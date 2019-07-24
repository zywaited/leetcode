package _127

import (
	"testing"

	"github.com/stretchr/testify/require"

	"leetcode/101_150/127/one"
)

func TestLadderLength(t *testing.T) {
	require.Equal(t, 5, one.LadderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
	require.Empty(t, 0, one.LadderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}))
}
