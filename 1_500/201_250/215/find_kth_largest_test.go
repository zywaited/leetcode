package _215

import (
	"leetcode/1_500/201_250/215/four"
	"leetcode/1_500/201_250/215/one"
	"leetcode/1_500/201_250/215/three"
	"leetcode/1_500/201_250/215/two"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindKthLargest(t *testing.T) {
	nums := [][]int{
		{3, 2, 1, 5, 6, 4},
		{3, 2, 3, 1, 2, 4, 5, 5, 6},
		{1},
	}
	ks := []int{2, 4, 1}
	outs := []int{5, 4, 1}
	for i, num := range nums {
		require.Equal(t, outs[i], one.FindKthLargest(num, ks[i]))
		require.Equal(t, outs[i], two.FindKthLargest(num, ks[i]))
		require.Equal(t, outs[i], three.FindKthLargest(num, ks[i]))
		require.Equal(t, outs[i], four.FindKthLargest(num, ks[i]))
	}
}
