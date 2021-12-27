## [实现 Trie (前缀树)](https://leetcode-cn.com/problems/implement-trie-prefix-tree/)

### 说明

* 实现一个 Trie (前缀树)，包含 insert, search, 和 startsWith 这三个操作。
* 你可以假设所有的输入都是由小写字母 a-z 构成的。
* 保证所有输入均为非空字符串。

### 实例

Trie trie = new Trie();

trie.insert("apple");
trie.search("apple");   // 返回 true
trie.search("app");     // 返回 false
trie.startsWith("app"); // 返回 true
trie.insert("app");   
trie.search("app");     // 返回 true

### 实现
需要细微的调节，Constructor返回的指针为了方便些测试用例，题目是返回实体
* one 本题的实现
* another 另一种前缀树(压缩树的高度)，基于字典树实现的前缀匹配，目前GO WEB框架echo的路由就是这样实现的，所以在做本题的时候顺带实现了
