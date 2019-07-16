package _231

import (
	"testing"

	"leetcode/201_250/231/one"

	"github.com/stretchr/testify/require"
)

func TestIsPowerOfTwo(t *testing.T) {
	require.NotEmpty(t, one.IsPowerOfTwo(1))
	require.NotEmpty(t, one.IsPowerOfTwo(16))
	require.Empty(t, one.IsPowerOfTwo(218))
}
