## [318. 最大单词长度乘积](https://leetcode-cn.com/problems/maximum-product-of-word-lengths/)

### 说明
给定一个字符串数组 words，找到 length(word[i]) * length(word[j]) 的最大值，并且这两个单词不含有公共字母。
你可以认为每个单词只包含小写字母。如果不存在这样的两个单词，返回 0。

* 2 <= words.length <= 1000
* 1 <= words[i].length <= 1000
* words[i] 仅包含小写字母

### 实例
#### 1
输入: ["abcw","baz","foo","bar","xtfn","abcdef"]
输出: 16

#### 2
输入: ["a","ab","abc","d","cd","bcd","abcd"]
输出: 4