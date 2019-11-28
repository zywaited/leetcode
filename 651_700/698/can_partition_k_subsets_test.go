package _698

import (
	"testing"

	"leetcode/651_700/698/one"
	"leetcode/651_700/698/two"

	"github.com/stretchr/testify/require"
)

func TestCanPartitionKSubsets(t *testing.T) {
	require.Equal(
		t,
		one.CanPartitionKSubsets([]int{4, 3, 2, 3, 5, 2, 1}, 4),
		two.CanPartitionKSubsets([]int{4, 3, 2, 3, 5, 2, 1}, 4),
	)

	require.Equal(
		t,
		one.CanPartitionKSubsets([]int{5, 5, 5, 5}, 4),
		two.CanPartitionKSubsets([]int{5, 5, 5, 5}, 4),
	)
}
