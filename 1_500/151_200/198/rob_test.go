package _198

import (
	"leetcode/1_500/151_200/198/one"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRob(t *testing.T) {
	require.Equal(t, 40, one.Rob([]int{1, 2, 3, 1, 5, 3, 5, 10, 9, 11, 4, 2, 4, 8}))
}
