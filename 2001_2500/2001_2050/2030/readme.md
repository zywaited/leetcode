## [2030. 含特定字母的最小子序列](https://leetcode-cn.com/problems/smallest-k-length-subsequence-with-occurrences-of-a-letter/)

### 说明
给你一个字符串 s ，一个整数 k ，一个字母 letter 以及另一个整数 repetition 。

返回 s 中长度为 k 且 字典序最小 的子序列，该子序列同时应满足字母 letter 出现 至少 repetition 次。
生成的测试用例满足 letter 在 s 中出现 至少 repetition 次。

子序列 是由原字符串删除一些（或不删除）字符且不改变剩余字符顺序得到的剩余字符串。

字符串 a 字典序比字符串 b 小的定义为：在 a 和 b 出现不同字符的第一个位置上，字符串 a 的字符在字母表中的顺序早于字符串 b 的字符。

提示
* 1 <= repetition <= k <= s.length <= 5 * 10^4
* s 由小写英文字母组成
* letter 是一个小写英文字母，在 s 中至少出现 repetition 次

### 实例
#### 1
输入：s = "leetcode", k = 4, letter = "e", repetition = 2
输出："ecde"

#### 2
输入：s = "bb", k = 2, letter = "b", repetition = 2
输出："bb"