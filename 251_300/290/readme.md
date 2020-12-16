## [290. 单词规律](https://leetcode-cn.com/problems/word-pattern/)

### 说明
给定一种规律 pattern 和一个字符串 str ，判断 str 是否遵循相同的规律。

这里的 遵循 指完全匹配，例如， pattern 里的每个字母和字符串 str 中的每个非空单词之间存在着双向连接的对应规律。

* 你可以假设 pattern 只包含小写字母， str 包含了由单个空格分隔的小写字母。

### 实例
#### 1
输入: pattern = "abba", str = "dog cat cat dog"
输出: true

#### 2
输入:pattern = "abba", str = "dog cat cat fish"
输出: false