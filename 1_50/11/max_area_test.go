package one

import (
	"testing"

	"leetcode/1_50/11/one"

	"github.com/stretchr/testify/require"
)

func TestMaxArea(t *testing.T) {
	require.Equal(t, 49, one.MaxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
