package _198

import (
	"testing"

	"github.com/stretchr/testify/require"

	"leetcode/151_200/198/one"
)

func TestRob(t *testing.T) {
	require.Equal(t, 40, one.Rob([]int{1, 2, 3, 1, 5, 3, 5, 10, 9, 11, 4, 2, 4, 8}))
}
