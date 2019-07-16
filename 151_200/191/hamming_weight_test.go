package _191

import (
	"testing"

	"leetcode/151_200/191/one"

	"github.com/stretchr/testify/require"
)

func TestHammingWeight(t *testing.T) {
	require.Equal(t, 3, one.HammingWeight(11))
	require.Equal(t, 1, one.HammingWeight(128))
	require.Equal(t, 31, one.HammingWeight(1<<32-2))
}
