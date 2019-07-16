package _208

import (
	"testing"

	"leetcode/201_250/208/another"
	"leetcode/201_250/208/one"

	"github.com/stretchr/testify/require"
)

type trie interface {
	Insert(string)

	Search(string) bool

	StartsWith(string) bool
}

func test(t *testing.T, tr trie) {
	tr.Insert("apple")
	require.NotEmpty(t, tr.Search("apple"))
	require.Empty(t, tr.Search("app"))
	require.NotEmpty(t, tr.StartsWith("app"))
	tr.Insert("app")
	require.NotEmpty(t, tr.Search("app"))
}

func TestTrie(t *testing.T) {
	test(t, one.Constructor())
	test(t, another.Constructor())
}
