package _990

import (
	"leetcode/501_1000/951_1000/990/one"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEquationsPossible(t *testing.T) {
	require.Empty(t, one.EquationsPossible([]string{"a==b", "b!=a"}))
	require.NotEmpty(t, one.EquationsPossible([]string{"b==a", "a==b"}))
	require.NotEmpty(t, one.EquationsPossible([]string{"a==b", "b==c", "a==c"}))
	require.Empty(t, one.EquationsPossible([]string{"a==b", "b!=c", "c==a"}))
	require.NotEmpty(t, one.EquationsPossible([]string{"c==c", "b==d", "x!=z"}))
}
