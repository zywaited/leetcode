package _342

import (
	"testing"

	"leetcode/301_350/342/one"

	"github.com/stretchr/testify/require"
)

func TestIsPowerOfFour(t *testing.T) {
	require.NotEmpty(t, one.IsPowerOfFour(16))
	require.Empty(t, one.IsPowerOfFour(32))
}
