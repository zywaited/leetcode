package _29

import (
	"testing"

	"leetcode/1_50/29/one"

	"github.com/stretchr/testify/require"
)

func TestDivisor(t *testing.T) {
	require.Equal(t, 3, one.Divide(10, 3))
	require.Equal(t, -2, one.Divide(7, -3))
}
