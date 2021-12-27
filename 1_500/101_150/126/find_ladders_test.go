package _126

import (
	"leetcode/1_500/101_150/126/one"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindLadders(t *testing.T) {
	outs := [][]string{
		{"hit", "hot", "dot", "dog", "cog"},
		{"hit", "hot", "lot", "log", "cog"},
	}
	finds := one.FindLadders("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"})
	require.Len(t, finds, len(outs))
	for _, out := range outs {
		require.Contains(t, finds, out)
	}

	outs = [][]string{
		{"a", "c"},
	}
	finds = one.FindLadders("a", "c", []string{"a", "b", "c"})
	require.Len(t, finds, len(outs))
	for _, out := range outs {
		require.Contains(t, finds, out)
	}
}
