package _347

import (
	"leetcode/301_350/347/one"
	"leetcode/301_350/347/two"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTopKFrequent(t *testing.T) {
	require.Equal(
		t,
		one.TopKFrequent([]int{1, 1, 1, 2, 2, 3}, 2),
		two.TopKFrequent([]int{1, 1, 1, 2, 2, 3}, 2),
	)
	require.Equal(
		t,
		one.TopKFrequent([]int{1}, 1),
		two.TopKFrequent([]int{1}, 1),
	)
}
