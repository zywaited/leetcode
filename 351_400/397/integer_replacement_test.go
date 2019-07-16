package _397

import (
	"math"
	"testing"

	"leetcode/351_400/397/one"
	"leetcode/351_400/397/two"

	"github.com/stretchr/testify/require"
)

type integerReplacement func(int) int

func test(t *testing.T, fn integerReplacement) {
	require.EqualValues(t, 3, fn(8))
	require.EqualValues(t, 4, fn(7))
	require.EqualValues(t, 17, fn(65535))
	require.EqualValues(t, 32, fn(math.MaxInt32))
}

func TestIntegerReplacement(t *testing.T) {
	test(t, one.IntegerReplacement)
	test(t, two.IntegerReplacement)
}
