package _1094

import (
	"leetcode/1001_1500/1051_1100/1094/one"
	"leetcode/1001_1500/1051_1100/1094/two"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCarPooling(t *testing.T) {
	require.Empty(t, one.CarPooling([][]int{{2, 1, 5}, {3, 3, 7}}, 4))
	require.NotEmpty(t, one.CarPooling([][]int{{2, 1, 5}, {3, 3, 7}}, 5))
	require.NotEmpty(t, one.CarPooling([][]int{{2, 1, 5}, {3, 5, 7}}, 3))
	require.NotEmpty(t, one.CarPooling([][]int{{3, 2, 7}, {3, 7, 9}, {8, 3, 9}}, 11))
	require.Empty(
		t,
		one.CarPooling([][]int{{9, 3, 6}, {8, 1, 7}, {6, 6, 8}, {8, 4, 9}, {4, 2, 9}}, 28),
	)
	require.NotEmpty(
		t,
		one.CarPooling(
			[][]int{{2, 4, 9}, {5, 8, 9}, {6, 5, 8}, {10, 2, 8}, {9, 1, 9}, {8, 5, 9}},
			35,
		),
	)

	require.Empty(t, two.CarPooling([][]int{{2, 1, 5}, {3, 3, 7}}, 4))
	require.NotEmpty(t, two.CarPooling([][]int{{2, 1, 5}, {3, 3, 7}}, 5))
	require.NotEmpty(t, two.CarPooling([][]int{{2, 1, 5}, {3, 5, 7}}, 3))
	require.NotEmpty(t, two.CarPooling([][]int{{3, 2, 7}, {3, 7, 9}, {8, 3, 9}}, 11))
	require.Empty(
		t,
		two.CarPooling([][]int{{9, 3, 6}, {8, 1, 7}, {6, 6, 8}, {8, 4, 9}, {4, 2, 9}}, 28),
	)
	require.NotEmpty(
		t,
		two.CarPooling(
			[][]int{{2, 4, 9}, {5, 8, 9}, {6, 5, 8}, {10, 2, 8}, {9, 1, 9}, {8, 5, 9}},
			35,
		),
	)
}
