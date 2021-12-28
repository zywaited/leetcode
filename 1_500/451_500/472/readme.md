## [472. 连接词](https://leetcode-cn.com/problems/concatenated-words/)

### 说明
给你一个 不含重复 单词的字符串数组 words ，请你找出并返回 words 中的所有 连接词 。

连接词 定义为：一个完全由给定数组中的至少两个较短单词组成的字符串。

提示：
* 1 <= words.length <= 10⁴
* 0 <= words[i].length <= 1000
* words[i] 仅由小写字母组成
* 0 <= sum(words[i].length) <= 10⁵

### 实例
#### 1
输入：words = ["cat","cats","catsdogcats","dog","dogcatsdog","hippopotamuses","rat","ratcatdogcat"]
输出：["catsdogcats","dogcatsdog","ratcatdogcat"]

#### 2
输入：words = ["cat","dog","catdog"]
输出：["catdog"]