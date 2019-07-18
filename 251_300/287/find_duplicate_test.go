package _287

import (
	"testing"

	"github.com/stretchr/testify/require"

	"leetcode/251_300/287/one"
	"leetcode/251_300/287/two"
)

type findDuplicate func([]int) int

func test(t *testing.T, fn findDuplicate) {
	require.Equal(t, 2, fn([]int{1, 3, 4, 2, 2}))
}

func TestFindDuplicate(t *testing.T) {
	test(t, one.FindDuplicate)
	test(t, two.FindDuplicate)
}
