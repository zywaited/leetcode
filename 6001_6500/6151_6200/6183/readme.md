## [6183. 字符串的前缀分数和](https://leetcode.cn/problems/sum-of-prefix-scores-of-strings/)

### 说明
给你一个长度为 n 的数组 words ，该数组由 非空 字符串组成。

定义字符串 word 的 分数 等于以 word 作为 前缀 的 words[i] 的数目。

例如，如果 words = ["a", "ab", "abc", "cab"] ，那么 "ab" 的分数是 2 ，因为 "ab" 是 "ab" 和 "abc" 的一个前缀。
返回一个长度为 n 的数组 answer ，其中 answer[i] 是 words[i] 的每个非空前缀的分数 总和 。

注意：字符串视作它自身的一个前缀。

### 提示：
* 1 <= words.length <= 1000
* 1 <= words[i].length <= 1000
* words[i] 由小写英文字母组成

### 实例
#### 1
输入：words = ["abc","ab","bc","b"]
输出：[5,4,3,2]

#### 2
输入：words = ["abcd"]
输出：[4]