package _337

import (
	"testing"

	"github.com/stretchr/testify/require"

	"leetcode/301_350/337/one"
)

func TestRob(t *testing.T) {
	// [3,2,3,null,3,null,1]
	in := &one.TreeNode{
		Val: 3,
		Left: &one.TreeNode{
			Val: 2,
			Right: &one.TreeNode{
				Val: 3,
			},
		},
		Right: &one.TreeNode{
			Val: 3,
			Right: &one.TreeNode{
				Val: 1,
			},
		},
	}

	require.Equal(t, 7, one.Rob(in))
	// [2,1,3,null,4]
	in = &one.TreeNode{
		Val: 2,
		Left: &one.TreeNode{
			Val: 1,
			Right: &one.TreeNode{
				Val: 4,
			},
		},
		Right: &one.TreeNode{
			Val: 3,
		},
	}

	require.Equal(t, 7, one.Rob(in))
}
