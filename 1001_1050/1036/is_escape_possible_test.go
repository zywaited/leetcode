package _1036

import (
	"testing"

	"leetcode/1001_1050/1036/one"

	"github.com/stretchr/testify/require"
)

func TestIsEscapePossible(t *testing.T) {
	require.NotEmpty(t, one.IsEscapePossible(
		[][]int{{0, 3}, {1, 0}, {1, 1}, {1, 2}, {1, 3}},
		[]int{0, 0},
		[]int{0, 2},
	))
}
