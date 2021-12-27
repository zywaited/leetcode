package _150

import (
	"leetcode/1_500/101_150/150/one"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEvalRPN(t *testing.T) {
	require.Equal(t, 9, one.EvalRPN([]string{"2", "1", "+", "3", "*"}))
	require.Equal(t, 6, one.EvalRPN([]string{"4", "13", "5", "/", "+"}))
	require.Equal(t, 22, one.EvalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))
}
