package _127

import (
	"leetcode/1_500/101_150/127/one"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLadderLength(t *testing.T) {
	require.Equal(t, 5, one.LadderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
	require.Empty(t, 0, one.LadderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}))
}
