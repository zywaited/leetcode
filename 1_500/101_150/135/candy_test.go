package _135

import (
	"leetcode/1_500/101_150/135/one"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCandy(t *testing.T) {
	require.Equal(t, 5, one.Candy([]int{1, 0, 5}))
	require.Equal(t, 4, one.Candy([]int{1, 2, 2}))
	require.Equal(t, 27, one.Candy([]int{1, 2, 2, 1, 0, 2, 1, 3, 4, 1, 6, 1, 0, 5, 3}))
	require.Equal(t, 13, one.Candy([]int{0, 1, 2, 3, 2, 1}))
}
