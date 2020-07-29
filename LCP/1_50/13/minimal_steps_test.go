package _13

import (
	"testing"

	"github.com/stretchr/testify/require"

	"leetcode/LCP/1_50/13/one"
)

func TestCode(t *testing.T) {
	require.Equal(t, 16, one.MinimalSteps([]string{"S#O", "M..", "M.T"}))
	require.Equal(t, -1, one.MinimalSteps([]string{"S#O", "M.#", "M.T"}))
	require.Equal(t, 17, one.MinimalSteps([]string{"S#O", "M.T", "M.."}))
	require.Equal(t, 9, one.MinimalSteps([]string{"T#O", ".##", "O..", ".#.", "OSM", "#.."}))
}
