package _30

import (
	"leetcode/1_500/1_50/30/one"
	"leetcode/1_500/1_50/30/two"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindSubstring(t *testing.T) {
	needs := []int{0, 9}
	finds := one.FindSubstring("barfoothefoobarman", []string{"foo", "bar"})
	require.Len(t, finds, len(needs))
	for _, need := range needs {
		require.Contains(t, finds, need)
	}

	needs = []int{0}
	finds = two.FindSubstring("cccab", []string{"ca", "cc"})
	require.Len(t, finds, len(needs))
	for _, need := range needs {
		require.Contains(t, finds, need)
	}
}
