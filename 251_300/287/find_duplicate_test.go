package _287

import (
	"testing"

	"leetcode/251_300/287/one"

	"github.com/stretchr/testify/require"
)

func TestFindDuplicate(t *testing.T) {
	require.Equal(t, 2, one.FindDuplicate([]int{1, 3, 4, 2, 2}))
	require.Equal(t, 2, one.FindDuplicate([]int{1, 3, 4, 2, 2}))
}
