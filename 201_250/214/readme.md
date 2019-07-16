## [最短回文串](https://leetcode-cn.com/problems/shortest-palindrome/)
### 说明

给定一个字符串 s，你可以通过在字符串前面添加字符将其转换为回文串。找到并返回可以用这种方式转换的最短回文串。

### 实例
#### 1

输入: "aacecaaa"
输出: "aaacecaaa"

#### 2

输入: "abcd"
输出: "dcbabcd"

### 实现
* one 利用KMP的Next数组
* two 基于Manacher算法

### 扩展
本题是在前面添加字符，如果都在前面添加或者都在后面添加字符, 找出最短回文串