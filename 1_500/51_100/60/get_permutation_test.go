package _60

import (
	"leetcode/1_500/51_100/60/one"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetPermutation(t *testing.T) {
	require.Equal(t, "213", one.GetPermutation(3, 3))
	require.Equal(t, "21", one.GetPermutation(2, 2))
	require.Equal(t, "132", one.GetPermutation(3, 2))
}
