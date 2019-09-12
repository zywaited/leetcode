package _21

import (
	"leetcode/301_350/321/one"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaxNumber(t *testing.T) {
	require.Equal(
		t,
		[]int{9, 8, 6, 5, 3},
		one.MaxNumber(
			[]int{3, 4, 6, 5},
			[]int{9, 1, 2, 5, 8, 3},
			5,
		),
	)
	require.Equal(
		t,
		[]int{6, 7, 6, 0, 4},
		one.MaxNumber(
			[]int{6, 7},
			[]int{6, 0, 4},
			5,
		),
	)
	require.Equal(
		t,
		[]int{9, 8, 9},
		one.MaxNumber(
			[]int{3, 9},
			[]int{8, 9},
			3,
		),
	)
	require.Equal(
		t,
		[]int{9, 8, 9, 6, 6},
		one.MaxNumber(
			[]int{3, 9, 6},
			[]int{8, 9, 6},
			5,
		),
	)
	require.Equal(
		t,
		[]int{9, 9, 9, 9, 9, 8, 5, 8, 6, 3, 9, 8},
		one.MaxNumber(
			[]int{3, 4, 6, 5, 9, 1, 9, 9, 9, 8},
			[]int{9, 1, 2, 5, 8, 6, 3, 9, 8},
			12,
		),
	)
	require.Equal(
		t,
		[]int{9, 9, 9, 9, 9, 9, 9, 9, 8, 9, 8, 8, 8, 5, 1, 4, 7, 8, 6, 9},
		one.MaxNumber(
			[]int{3, 4, 6, 5, 9, 1, 9, 9, 9, 8, 9, 9, 8, 8, 5, 1, 4, 7, 8, 6, 9},
			[]int{9, 1, 2, 5, 8, 6, 3, 9, 8, 8, 5, 6, 7, 7, 8, 8, 9, 8},
			20,
		),
	)
	require.Equal(
		t,
		[]int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 8, 8, 6, 7, 7, 8, 9, 7, 7, 5, 5, 8, 9, 4, 8, 2, 8, 9, 3},
		one.MaxNumber(
			[]int{3, 4, 6, 5, 6, 1, 7, 8, 9, 1, 9, 9, 8, 8, 4, 2, 9, 8, 5, 1, 9, 8, 4, 8, 2, 8, 9, 3},
			[]int{9, 1, 2, 5, 8, 3, 9, 8, 8, 9, 4, 6, 7, 1, 2, 4, 6, 7, 8, 9, 9, 9, 8, 6, 7, 7, 8, 9, 7, 7, 5, 5, 8, 9},
			30,
		),
	)
}
