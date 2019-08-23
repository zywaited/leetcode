package _59

import (
	"testing"

	"github.com/stretchr/testify/require"

	"leetcode/51_100/59/one"
)

func TestGenerateMatrix(t *testing.T) {
	expects := [][][]int{
		{
			{1, 2, 3},
			{8, 9, 4},
			{7, 6, 5},
		},
		{
			{1, 2, 3, 4, 5},
			{16, 17, 18, 19, 6},
			{15, 24, 25, 20, 7},
			{14, 23, 22, 21, 8},
			{13, 12, 11, 10, 9},
		},
	}
	ins := []int{3, 5}
	for i, in := range ins {
		require.Equal(t, expects[i], one.GenerateMatrix(in))
	}
}
