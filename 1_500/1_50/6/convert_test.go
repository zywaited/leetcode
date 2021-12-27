package _6

import (
	"leetcode/1_500/1_50/6/one"
	"leetcode/1_500/1_50/6/two"
	"testing"

	"github.com/stretchr/testify/require"
)

type convert func(string, int) string

func test(t *testing.T, fn convert) {
	require.Equal(t, "LCIRETOESIIGEDHN", fn("LEETCODEISHIRING", 3))
	require.Equal(t, "LDREOEIIECIHNTSG", fn("LEETCODEISHIRING", 4))
}

func TestConvert(t *testing.T) {
	test(t, one.Convert)
	test(t, two.Convert)
}
