package _877

import (
	"testing"

	"leetcode/851_900/877/two"

	"github.com/stretchr/testify/require"
)

func TestStoneGame(t *testing.T) {
	require.NotEmpty(t, two.StoneGame([]int{5, 3, 4, 5}))
}
