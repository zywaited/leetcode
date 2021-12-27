package one

import (
	"leetcode/1_500/201_250/229/one"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMajorityElement(t *testing.T) {
	require.Equal(t, []int{3}, one.MajorityElement([]int{3, 2, 3}))
	require.Equal(t, []int{1, 2}, one.MajorityElement([]int{1, 1, 1, 3, 3, 2, 2, 2}))
}
