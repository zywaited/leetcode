package _137

import (
	"leetcode/1_500/101_150/137/one"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSingleNumber(t *testing.T) {
	require.Equal(t, 3, one.SingleNumber([]int{2, 2, 3, 2}))
	require.Equal(t, 99, one.SingleNumber([]int{0, 1, 0, 1, 0, 1, 99}))
}
