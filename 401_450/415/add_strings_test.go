package _415

import (
	"testing"

	"github.com/stretchr/testify/require"

	"leetcode/401_450/415/one"
)

func TestAddStrings(t *testing.T) {
	require.Equal(t, "3", one.AddStrings("1", "2"))
	require.Equal(t, "", one.AddStrings("", ""))
	require.Equal(t, "11", one.AddStrings("9", "2"))
	require.Equal(t, "10098", one.AddStrings("99", "9999"))
}
