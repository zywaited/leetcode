package _389

import (
	"leetcode/1_500/351_400/389/one"
	"leetcode/1_500/351_400/389/two"
	"testing"

	"github.com/stretchr/testify/require"
)

type findTheDifference func(string, string) byte

func test(t *testing.T, fn findTheDifference) {
	require.Equal(t, byte('e'), fn("abbcad", "babcdae"))
}

func TestFindTheDifference(t *testing.T) {
	test(t, one.FindTheDifference)
	test(t, two.FindTheDifference)
}
