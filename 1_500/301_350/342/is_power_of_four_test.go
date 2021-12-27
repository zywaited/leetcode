package _342

import (
	"leetcode/1_500/301_350/342/one"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsPowerOfFour(t *testing.T) {
	require.NotEmpty(t, one.IsPowerOfFour(16))
	require.Empty(t, one.IsPowerOfFour(32))
}
