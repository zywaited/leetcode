package _30

import (
	"testing"

	"github.com/stretchr/testify/require"

	"leetcode/1_50/30/one"
	"leetcode/1_50/30/two"
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
