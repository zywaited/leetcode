package _231

import (
	"leetcode/1_500/201_250/231/one"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsPowerOfTwo(t *testing.T) {
	require.NotEmpty(t, one.IsPowerOfTwo(1))
	require.NotEmpty(t, one.IsPowerOfTwo(16))
	require.Empty(t, one.IsPowerOfTwo(218))
}
