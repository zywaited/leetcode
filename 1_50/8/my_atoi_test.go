package _8

import (
	"testing"

	"leetcode/1_50/8/one"

	"github.com/stretchr/testify/require"
)

func TestMyAtoi(t *testing.T) {
	require.Equal(t, 24, one.MyAtoi("24"))
	require.Equal(t, -42, one.MyAtoi("   -42"))
}
