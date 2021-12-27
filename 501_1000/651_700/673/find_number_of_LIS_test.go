package _673

import (
	"leetcode/501_1000/651_700/673/one"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindNumberOfLIS(t *testing.T) {
	require.Equal(t, 2, one.FindNumberOfLIS([]int{1, 3, 5, 4, 7}))
	require.Equal(t, 5, one.FindNumberOfLIS([]int{2, 2, 2, 2, 2}))
}
