package _268

import (
	"testing"

	"leetcode/251_300/268/one"
	"leetcode/251_300/268/two"

	"github.com/stretchr/testify/require"
)

type missingNumber func([]int) int

func test(t *testing.T, fn missingNumber) {
	require.Equal(t, 2, fn([]int{3, 0, 1}))
	require.Equal(t, 8, fn([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
}

func TestMissingNumber(t *testing.T) {
	test(t, one.MissingNumber)
	test(t, two.MissingNumber)
}
