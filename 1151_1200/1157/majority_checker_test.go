package _1157

import (
	"testing"

	"github.com/stretchr/testify/require"

	"leetcode/1151_1200/1157/one"
	"leetcode/1151_1200/1157/two"
)

func TestMajorityChecker(t *testing.T) {
	omc := one.Constructor([]int{1, 1, 2, 2, 1, 1})
	require.Equal(t, 1, omc.Query(0, 5, 4))
	require.Equal(t, -1, omc.Query(0, 3, 3))
	require.Equal(t, 2, omc.Query(2, 3, 2))

	tmc := two.Constructor([]int{1, 1, 2, 2, 1, 1})
	require.Equal(t, 1, tmc.Query(0, 5, 4))
	require.Equal(t, -1, tmc.Query(0, 3, 3))
	require.Equal(t, 2, tmc.Query(2, 3, 2))
}
