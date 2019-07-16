package _80

import (
	"testing"

	"leetcode/51_100/80/one"

	"github.com/stretchr/testify/require"
)

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 3}
	nl := one.RemoveDuplicates(nums)
	outs := []int{1, 1, 2, 2, 3}
	require.Equal(t, len(outs), nl)
	for i, out := range outs {
		require.Equal(t, out, nums[i])
	}
}
