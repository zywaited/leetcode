## [2516. 每种字符至少取 K 个](https:leetcode.cn/problems/take-k-of-each-character-from-left-and-right/)

### 说明
给你一个由字符 'a'、'b'、'c' 组成的字符串 s 和一个非负整数 k 。每分钟，你可以选择取走 s 最左侧 还是 最右侧 的那个字符。

你必须取走每种字符 至少 k 个，返回需要的 最少 分钟数；如果无法取到，则返回 -1 。

### 提示：
* 1 <= s.length <= 10^5
* 0 <= k <= s.length
* s 仅由字母 'a'、'b'、'c' 组成

### 实例
#### 1
输入：s = "aabaaaacaabc", k = 2
输出：8

#### 2
输入：s = "a", k = 1
输出：-1