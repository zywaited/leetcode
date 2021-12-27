## [字符串的排列](https://leetcode-cn.com/problems/permutation-in-string/)
### 说明

给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的排列。

换句话说，第一个字符串的排列之一是第二个字符串的子串。。

### 实例
#### 1

输入: s1 = "ab" s2 = "eidbaooo"
输出: True
解释: s2 包含 s1 的排列之一 ("ba").

### 实现
* one 滑动窗口
* two 位运算(计算得分)