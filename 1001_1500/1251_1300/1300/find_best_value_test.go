package _1300

import (
	"leetcode/1001_1500/1251_1300/1300/one"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindBestValue(t *testing.T) {
	require.EqualValues(t, 3, one.FindBestValue([]int{4, 9, 3}, 10))
	require.EqualValues(t, 5, one.FindBestValue([]int{2, 3, 5}, 10))
	require.EqualValues(t, 4, one.FindBestValue([]int{2, 3, 5}, 9))
	require.EqualValues(t, 6, one.FindBestValue([]int{1, 3, 6}, 10))
	require.EqualValues(t, 11361, one.FindBestValue([]int{60864, 25176, 27249, 21296, 20204}, 56803))
}
