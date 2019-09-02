package _1094

import (
	"testing"

	"github.com/stretchr/testify/require"

	"leetcode/1051_1100/1094/one"
	"leetcode/1051_1100/1094/two"
)

func TestCarPooling(t *testing.T) {
	require.Empty(t, one.CarPooling([][]int{{2, 1, 5}, {3, 3, 7}}, 4))
	require.NotEmpty(t, one.CarPooling([][]int{{2, 1, 5}, {3, 3, 7}}, 5))
	require.NotEmpty(t, one.CarPooling([][]int{{2, 1, 5}, {3, 5, 7}}, 3))
	require.NotEmpty(t, one.CarPooling([][]int{{3, 2, 7}, {3, 7, 9}, {8, 3, 9}}, 11))

	// 这种情况结果是正确的，但是力扣第一个算法通不过
	require.Equal(
		t,
		one.CarPooling([][]int{{9, 3, 6}, {8, 1, 7}, {6, 6, 8}, {8, 4, 9}, {4, 2, 9}}, 28),
		two.CarPooling([][]int{{9, 3, 6}, {8, 1, 7}, {6, 6, 8}, {8, 4, 9}, {4, 2, 9}}, 28),
	)
}
