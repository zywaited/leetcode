package _50

import (
	"leetcode/1_500/1_50/50/one"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMyPow(t *testing.T) {
	require.Equal(t, float64(1024), one.MyPow(float64(2), 10))
	require.Equal(t, float64(27), one.MyPow(float64(3), 3))
	require.Equal(t, float64(0.25), one.MyPow(float64(4), -1))
}
