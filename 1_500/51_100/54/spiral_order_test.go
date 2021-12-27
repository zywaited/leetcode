package _54

import (
	"leetcode/1_500/51_100/54/one"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSpiralOrder(t *testing.T) {
	ins := [][][]int{
		{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
		{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
		},
		{
			{1, 2, 3, 4, 5},
			{6, 7, 8, 9, 10},
			{11, 12, 13, 14, 15},
			{16, 17, 18, 19, 20},
			{21, 22, 23, 24, 25},
		},
	}
	expects := [][]int{
		{1, 2, 3, 6, 9, 8, 7, 4, 5},
		{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		{1, 2, 3, 4, 5, 10, 15, 20, 25, 24, 23, 22, 21, 16, 11, 6, 7, 8, 9, 14, 19, 18, 17, 12, 13},
	}

	for i, in := range ins {
		require.Equal(t, expects[i], one.SpiralOrder(in))
	}
}
