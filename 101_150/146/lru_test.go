package _146

import (
	"testing"

	"github.com/stretchr/testify/require"

	"leetcode/101_150/146/one"
	"leetcode/101_150/146/two"
)

type LRUCache interface {
	Get(int) int
	Put(int, int)
}

func test(t *testing.T, cache LRUCache) {
	cache.Put(1, 1)
	cache.Put(2, 2)
	require.Equal(t, 1, cache.Get(1))  // 返回  1
	cache.Put(3, 3)                    // 该操作会使得密钥 2 作废
	require.Equal(t, -1, cache.Get(2)) // 返回 -1 (未找到)
	cache.Put(4, 4)                    // 该操作会使得密钥 1 作废
	require.Equal(t, -1, cache.Get(1)) // 返回 -1 (未找到)
	require.Equal(t, 3, cache.Get(3))  // 返回  3
	require.Equal(t, 4, cache.Get(4))  // 返回  4
}

func TestLRUCache(t *testing.T) {
	cache := one.Constructor(2)
	test(t, &cache)
	cache2 := two.Constructor(2)
	test(t, &cache2)
}
