package _287

import (
	"leetcode/1_500/251_300/287/one"
	"leetcode/1_500/251_300/287/two"
	"testing"

	"github.com/stretchr/testify/require"
)

type findDuplicate func([]int) int

func test(t *testing.T, fn findDuplicate) {
	require.Equal(t, 2, fn([]int{1, 3, 4, 2, 2}))
}

func TestFindDuplicate(t *testing.T) {
	test(t, one.FindDuplicate)
	test(t, two.FindDuplicate)
}
