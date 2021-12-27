## [1717. 删除子字符串的最大得分](https://leetcode-cn.com/problems/maximum-score-from-removing-substrings/)

### 说明
给你一个字符串 s 和两个整数 x 和 y 。你可以执行下面两种操作任意次。

删除子字符串 "ab" 并得到 x 分。
比方说，从 "cabxbae" 删除 ab ，得到 "cxbae" 。
删除子字符串"ba" 并得到 y 分。
比方说，从 "cabxbae" 删除 ba ，得到 "cabxe" 。
请返回对 s 字符串执行上面操作若干次能得到的最大得分。

* 1 <= s.length <= 105
* 1 <= x, y <= 104
* s 只包含小写英文字母。

### 实例
#### 1
输入：s = "cdbcbbaaabab", x = 4, y = 5
输出：19

#### 2
输入：s = "aabbaaxybbaabb", x = 5, y = 4
输出：20