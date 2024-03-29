package _179

import (
	"leetcode/1_500/151_200/179/one"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLargestNumber(t *testing.T) {
	require.Equal(t, "210", one.LargestNumber([]int{10, 2}))
	require.Equal(t, "9534330", one.LargestNumber([]int{3, 30, 34, 5, 9}))
}
