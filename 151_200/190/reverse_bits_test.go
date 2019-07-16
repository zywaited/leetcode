package one

import (
	"testing"

	"leetcode/151_200/190/one"

	"github.com/stretchr/testify/require"
)

func TestReverseBits(t *testing.T) {
	require.EqualValues(t, 964176192, one.ReverseBits(43261596))
	require.EqualValues(t, 3221225471, one.ReverseBits(4294967293))
}
