package _389

import (
	"testing"

	"leetcode/351_400/389/one"
	"leetcode/351_400/389/two"

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
