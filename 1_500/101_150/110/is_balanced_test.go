package one

import (
	"leetcode/1_500/101_150/110/one"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsBalanced(t *testing.T) {
	require.NotEmpty(t, one.IsBalanced(&one.TreeNode{
		Val: 3,
		Left: &one.TreeNode{
			Val: 9,
		},
		Right: &one.TreeNode{
			Val: 20,
			Left: &one.TreeNode{
				Val: 15,
			},
			Right: &one.TreeNode{
				Val: 7,
			},
		},
	}))
}
