package _48

import (
	"testing"

	"github.com/stretchr/testify/require"

	"leetcode/1_50/48/one"
)

func TestRotate(t *testing.T) {
	out := [][]int{
		{13, 9, 5, 1},
		{14, 10, 6, 2},
		{15, 11, 7, 3},
		{16, 12, 8, 4},
	}
	in := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	one.Rotate(in)
	require.Equal(t, out, in)
}
