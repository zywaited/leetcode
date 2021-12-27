package _169

import (
	"leetcode/1_500/151_200/169/one"
	"leetcode/1_500/151_200/169/two"
	"testing"

	"github.com/stretchr/testify/require"
)

type majorityElement func([]int) int

func test(t *testing.T, fn majorityElement) {
	require.Equal(t, 3, fn([]int{3, 2, 3}))
	require.Equal(t, 2, fn([]int{2, 2, 1, 1, 1, 2, 2}))
}

func TestMajorityElement(t *testing.T) {
	test(t, one.MajorityElement)
	test(t, two.MajorityElement)
}
