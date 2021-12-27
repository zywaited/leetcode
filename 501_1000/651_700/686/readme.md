## [686. 重复叠加字符串匹配](https://leetcode-cn.com/problems/repeated-string-match/)

### 说明
给定两个字符串 a 和 b，寻找重复叠加字符串 a 的最小次数，使得字符串 b 成为叠加后的字符串 a 的子串，如果不存在则返回 -1。

注意：字符串 "abc" 重复叠加 0 次是 ""，重复叠加 1 次是 "abc"，重复叠加 2 次是 "abcabc"。

提示：
* 1 <= a.length <= 10^4
* 1 <= b.length <= 10^4
* a 和 b 由小写英文字母组成

### 实例
#### 1
输入：a = "abcd", b = "cdabcdab"
输出：3

#### 2
输入：a = "a", b = "aa"
输出：2