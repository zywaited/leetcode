package _224

import (
	"leetcode/1_500/201_250/224/one"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalculator(t *testing.T) {
	require.Equal(t, 2, one.Calculator("1 + 1"))
	require.Equal(t, 3, one.Calculator(" 2-1 + 2 "))
	require.Equal(t, 23, one.Calculator("(1+(4+5+2)-3)+(6+8)"))
}
