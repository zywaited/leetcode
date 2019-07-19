## [LRU缓存机制](https://leetcode-cn.com/problems/lru-cache/)
### 说明

运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制。它应该支持以下操作： 获取数据 get 和 写入数据 put 。

获取数据 get(key) - 如果密钥 (key) 存在于缓存中，则获取密钥的值（总是正数），否则返回 -1。
写入数据 put(key, value) - 如果密钥不存在，则写入其数据值。当缓存容量达到上限时，它应该在写入新数据之前删除最近最少使用的数据值，从而为新的数据值留出空间。


### 实例
#### 1

LRUCache cache = new LRUCache( 2 /* 缓存容量 */ );

cache.put(1, 1);
cache.put(2, 2);
cache.get(1);       // 返回  1
cache.put(3, 3);    // 该操作会使得密钥 2 作废
cache.get(2);       // 返回 -1 (未找到)
cache.put(4, 4);    // 该操作会使得密钥 1 作废
cache.get(1);       // 返回 -1 (未找到)
cache.get(3);       // 返回  3
cache.get(4);       // 返回  4

### 实现
* one 保持插入顺序的map
(类似PHP的HASH TABLE(PHP是放在一个数组中的，可以参考: [PHP HASH TABLE](https://github.com/zywaited/RedisPool/blob/master/src/map/hash.h))，也就是索引和数据单独存放，这样索引是无序的，但数据是顺序插入)
这里为了方便就存成一个hash一个数组，数组频繁变动效率很慢
* two hash map + 双向链表(该数据结构删除和添加都是O(1)，访问和添加都把该元素放到链首，删除时就是链尾)